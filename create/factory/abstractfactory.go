package factory

type (
	DAOFactory interface {
		CreateOrderMainDAO() OrderMainDAO
		CreateOrderDetailDAO() OrderDetailDAO
	}
	OrderMainDAO   interface{ SaveOrderMain() }
	OrderDetailDAO interface{ SaveOrderDetail() }
)

type (
	TxtDAOFactory  struct{}
	JSONDAOFactory struct{}
	RDBDAOFactory  struct{}
)

type (
	TxtMainDao    struct{}
	TxtDetailDao  struct{}
	RDBMainDao    struct{}
	RDBDetailDao  struct{}
	JSONMainDao   struct{}
	JSONDetailDao struct{}
)

func (*TxtDAOFactory) CreateOrderMainDAO() OrderMainDAO      { return &TxtMainDao{} }
func (*TxtDAOFactory) CreateOrderDetailDAO() OrderDetailDAO  { return &TxtDetailDao{} }
func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO      { return &RDBMainDao{} }
func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO  { return &RDBDetailDao{} }
func (*JSONDAOFactory) CreateOrderMainDAO() OrderMainDAO     { return &JSONMainDao{} }
func (*JSONDAOFactory) CreateOrderDetailDAO() OrderDetailDAO { return &JSONDetailDao{} }

func (*TxtMainDao) SaveOrderMain()      {}
func (*TxtDetailDao) SaveOrderDetail()  {}
func (*RDBMainDao) SaveOrderMain()      {}
func (*RDBDetailDao) SaveOrderDetail()  {}
func (*JSONMainDao) SaveOrderMain()     {}
func (*JSONDetailDao) SaveOrderDetail() {}
