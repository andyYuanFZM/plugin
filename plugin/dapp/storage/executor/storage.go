package executor

import (
	log "github.com/33cn/chain33/common/log/log15"
	drivers "github.com/33cn/chain33/system/dapp"
	"github.com/33cn/chain33/types"
	storagetypes "github.com/33cn/plugin/plugin/dapp/storage/types"
	ac "github.com/33cn/plugin/plugin/dapp/accountmanager/executor"
	et "github.com/33cn/plugin/plugin/dapp/accountmanager/types"
)

/*
 * 执行器相关定义
 * 重载基类相关接口
 */

var (
	//日志
	elog = log.New("module", "storage.executor")
)

var driverName = storagetypes.StorageX

// Init register dapp
func Init(name string, cfg *types.Chain33Config, sub []byte) {
	drivers.Register(cfg, GetName(), newStorage, cfg.GetDappFork(driverName, "Enable"))
	InitExecType()
}

// InitExecType Init Exec Type
func InitExecType() {
	ety := types.LoadExecutorType(driverName)
	ety.InitFuncList(types.ListMethod(&storage{}))
}

type storage struct {
	drivers.DriverBase
}

func newStorage() drivers.Driver {
	t := &storage{}
	t.SetChild(t)
	t.SetExecutorType(types.LoadExecutorType(driverName))
	return t
}

// GetName get driver name
func GetName() string {
	return newStorage().GetName()
}

func (s *storage) GetDriverName() string {
	return driverName
}

//ExecutorOrder Exec 的时候 同时执行 ExecLocal
func (s *storage) ExecutorOrder() int64 {
	cfg := s.GetAPI().GetConfig()
	if cfg.IsDappFork(s.GetHeight(), storagetypes.StorageX, storagetypes.ForkStorageLocalDB) {
		return drivers.ExecLocalSameTime
	}
	return s.DriverBase.ExecutorOrder()
}

// CheckTx 实现自定义检验交易接口，供框架调用
func (s *storage) CheckTx(tx *types.Transaction, index int) error {

	var storage storagetypes.StorageAction
	types.Decode(tx.GetPayload(), &storage)

	from, _ := tx.GetViewFromToAddr()
	elog.Error("=======================from address", from)
	status, level ,_:= ac.GetStatusAndLevel(s.GetLocalDB(), from)
	elog.Error("=======================status", status, "level", level)
	// 冻结状态的账户，无法进行下一步交易
	if (status != et.Normal) {
		// 冻结或其它状态
		return et.ErrAccountIsFrozen
	}

	if storage.Ty == storagetypes.TyEncryptStorageAction {
		if (level <= 0) {
			// 普通权限不能操作
			return et.ErrAccountIsLowLevel
		}
	}
	return nil
}
