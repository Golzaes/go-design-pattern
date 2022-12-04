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
	RDBDAOFactory  struct{}
	JSONDAOFactory struct{}
)

type (
	RDBMainDAO    struct{}
	RDBDetailDAO  struct{}
	JSONMainDAO   struct{}
	JSONDetailDAO struct{}
)

func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO      { return &RDBMainDAO{} }
func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO  { return &RDBDetailDAO{} }
func (*JSONDAOFactory) CreateOrderMainDAO() OrderMainDAO     { return &JSONMainDAO{} }
func (*JSONDAOFactory) CreateOrderDetailDAO() OrderDetailDAO { return &JSONDetailDAO{} }

func (*RDBMainDAO) SaveOrderMain()      {}
func (*RDBDetailDAO) SaveOrderDetail()  {}
func (*JSONMainDAO) SaveOrderMain()     {}
func (*JSONDetailDAO) SaveOrderDetail() {}
