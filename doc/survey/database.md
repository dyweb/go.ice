# Database

## Config

### Laravel

- https://laravel.com/docs/5.5/database
- https://github.com/laravel/laravel/blob/master/config/database.php

````php
[
  'default' => env('DB_CONNECTION', 'mysql'),
  'connections' => [
    'sqlite' => [
      'driver' => 'sqlite',
      'database' => env('DB_DATABASE', database_path('database.sqlite')),
      'prefix' => '',
    ],
    'mysql' => [
      'read' => ['host' => '192.168.1.1'],
      'write' => ['host' => '196.168.1.12'],
      'stick' => true, // use the write connection for read in same request if there are previous write
      'driver' => 'mysql',
      'database' => 'foo',
      'username' => 'root',
      'password' => '',
      'charset' => 'utf8mb4',
      'collation' => 'utf8mb4_unicode_ci',
      'prefix' => '',
    ]
  ]
]
````

````php
$users = DB::connection('foo')->select(...);
$pdo = DB::connection()->getPdo();
DB::insert('insert into users (id, name) values (?, ?)', [1, 'Dayle']);
$users = DB::select('select * from users where active = ?', [1]);
$results = DB::select('select * from users where id = :id', ['id' => 1]);
$affected = DB::update('update users set votes = 100 where name = ?', ['John']);
$deleted = DB::delete('delete from users');
DB::statement('drop table users');
DB::transaction(function () {
    DB::table('users')->update(['votes' => 1]);
    DB::table('posts')->delete();
});
DB::beginTransaction();
DB::rollBack();
DB::commit();
````

### Ruby on Rails

- http://guides.rubyonrails.org/configuring.html 3.1.4 Configurating a database
- `pool` limit max number of concurrent connections to database

````yaml
# config/database.yml
development:
  adapter: sqlite3
  database: db/development.sqlite3
  pool: 5 # max number of connnection to database https://devcenter.heroku.com/articles/concurrency-and-database-connections#connection-pool
  timeout: 5000
development:
  adapter: mysql2
  encoding: utf8
  database: blog_development
  pool: 5
  username: root
  password:
  socket: /tmp/mysql.sock
development:
  adapter: postgresql
  encoding: unicode
  database: blog_development
  prepared_statements: false
  statement_limit: 200
````

### Django

- https://docs.djangoproject.com/en/2.0/ref/databases/
- https://docs.djangoproject.com/en/2.0/ref/settings/#databases

````python
DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.postgresql',
        'NAME': 'mydatabase',
        'USER': 'mydatabaseuser',
        'PASSWORD': 'mypassword',
        'HOST': '127.0.0.1',
        'PORT': '5432',
        'TEST': {
           'NAME': 'mytestdatabase',
        },
    }
}
````

### Flask

````python
# Load default config and override config from an environment variable
app.config.update(dict(
    DATABASE=os.path.join(app.root_path, 'flaskr.db'),
    SECRET_KEY='development key',
    USERNAME='admin',
    PASSWORD='default'
))
````

### Spring

````text
spring.datasource.url=jdbc:mysql://localhost/test
spring.datasource.username=dbuser
spring.datasource.password=dbpass
spring.datasource.driver-class-name=com.mysql.jdbc.Driver
````

- https://docs.spring.io/spring-boot/docs/current/reference/html/boot-features-sql.html
- https://stackoverflow.com/questions/11881548/jpa-or-jdbc-how-are-they-different
  - `JDBC` is a standard for database access
  - `JPA` is a standard for ORM 

### Dropwizard

- it is using a new interface called jdbi ...
- http://www.dropwizard.io/0.7.1/docs/manual/jdbi.html#configuration

````yaml
database:
  # the name of your JDBC driver
  driverClass: org.postgresql.Driver
  # the username
  user: pg-user
  # the password
  password: iAMs00perSecrEET
  # the JDBC URL
  url: jdbc:postgresql://db.example.com/db-prod
  # any properties specific to your JDBC driver:
  properties:
    charSet: UTF-8
  # the maximum amount of time to wait on an empty pool before throwing an exception
  maxWaitForConnection: 1s
  # the SQL query to run when validating a connection's liveness
  validationQuery: "/* MyService Health Check */ SELECT 1"
  # the minimum number of connections to keep open
  minSize: 8
  # the maximum number of connections to keep open
  maxSize: 32
  # whether or not idle connections should be validated
  checkConnectionWhileIdle: false
  # the amount of time to sleep between runs of the idle connection validation, abandoned cleaner and idle pool resizing
  evictionInterval: 10s
  # the minimum amount of time an connection must sit idle in the pool before it is eligible for eviction
  minIdleTime: 1 minute
````
Java Hibernate

Java Dropwizard