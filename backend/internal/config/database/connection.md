

# Configurar la db
dsn := "<root>:<password>@tcp(127.0.0.1:3306)/<nombre_database>?parseTime=true"

# Configuracion estandar si tu usuario no tiene contraseña
dsn := "root:@tcp(127.0.0.1:3306)/?parseTime=true"

    <nombre_database>: no se especifica
    root: usuario
    password: vacia, es decir, no tiene

