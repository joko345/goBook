go mod init github.com/joko345/goBook
go get github.com/jinzhu/gorm/dialects/mysql@latest
go get github.com/gorilla/mux
go mod tidy //memastikan semua dependecies terinstall

cd cmd
cd main
CD book-crud
npm start