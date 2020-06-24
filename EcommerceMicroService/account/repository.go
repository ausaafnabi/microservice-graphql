type Repository interface {
  Close()
  PutAccount(ctx context.Context,a Account) error
  GetAccountByID(ctx context.Context,id string)(*Account,error)
  ListAccounts(ctx context.Context,skip unit64 ,take uint64)([]Account ,error)
}

type postgresRepository struct {
  db *sql.DB
}

func NewPostgresRepository(url string )(Repository ,error){
  db,err :=sql.Open("postgres",url)
  if err != nil {
    return nil,err
  }
  error = db.Ping()
  if err != nil {
    return nil,err 
  }
  return &postgresRepository(db),nil 
}

