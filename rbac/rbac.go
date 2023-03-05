package rbac

import (
	"go-app/config"
	"go-app/lib/logger"
	"sync"

	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
)

var enforcer *casbin.Enforcer
var once sync.Once

func New() *casbin.Enforcer {
	once.Do(func() {
		var m casbinmodel.Model
		if config.CASBIN.Model == nil {
			m = casbinmodel.NewModel()
			m.AddDef("r", "r", "sub, obj, act")
			m.AddDef("p", "p", "sub, obj, act")
			m.AddDef("g", "g", "_, _")
			m.AddDef("e", "e", "some(where (p.eft == allow))")
			m.AddDef("m", "m", `g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act || checkSuperAdmin(r.sub)`)
		} else {
			m = config.CASBIN.Model
		}

		orm, err := gormadapter.NewAdapter(config.CASBIN.DriverName, config.CASBIN.DataSourceName, true)
		if err != nil {
			logger.Error("[Casbin]", zap.Error(err))
			return
		}
		enforcer, err = casbin.NewEnforcer(m, orm)
		if err != nil {
			logger.Error("[Casbin]", zap.Error(err))
			return
		}

		enforcer.AddFunction("checkSuperAdmin", func(arguments ...any) (any, error) {
			// 获取用户名
			username := arguments[0].(string)
			// 检查用户名的角色是否为superadmin
			return enforcer.HasRoleForUser(username, "superadmin")
		})
	})
	return enforcer
}

// 隐式判断一个用户是否有某个资源的操作权限
func HasImplicitPermissionsForUser(user, resource, permission string) bool {
	permissions, err := New().GetImplicitPermissionsForUser(user)
	if err != nil {
		return false
	}
	for _, p := range permissions {
		if p[1] == resource && p[2] == permission {
			return true
		}
	}
	return false
}
