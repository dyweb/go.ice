# Migration

## Go

- https://github.com/avelino/awesome-go#database

mattes/migrate

- https://github.com/mattes/migrate 1.9k star
  - read SQL file and go up and down, support using multiple file source https://github.com/mattes/migrate/tree/master/source
  - https://github.com/mattes/migrate/blob/master/MIGRATIONS.md use sql w/ `up` and `down` in file name

buffalo pop/soda

- https://github.com/markbates/pop/tree/master/soda 476 star, part of buffalo framework
  - https://gobuffalo.io/docs/db/migrations
  - https://github.com/markbates/pop/blob/master/fizz%2FREADME.md#a-common-dsl-for-migrating-databases
    - just a table builder (like what CI has)
  - https://github.com/markbates/pop/blob/master/schema_migrations.go#L7:5 `version (string & index)`
  - `migrate Up` https://github.com/markbates/pop/blob/master/migrator.go#L42:19
    - **use transaction for each migration** https://github.com/markbates/pop/blob/master/migrator.go#L55 
- migrator
  - file [file_migrator.go](https://github.com/markbates/pop/blob/master/file_migrator.go#L23:6)
    - walk the directory, read the content and execute them as sql ...
  - file embed in binary [migration_box.go](https://github.com/markbates/pop/blob/master/migration_box.go)
    - use https://github.com/gobuffalo/packr to embed static files into go binary, including migration files

sql-migrate

- https://github.com/rubenv/sql-migrate
  - mysql need `?parseTime=true` to use `time.Time`

## Django

- https://docs.djangoproject.com/en/2.0/topics/migrations/#migration-files
  - `dependencies = [('migrations', '0001_initial')]`
    - [ ] what if there are cycle in dependencies? it seems django will detect it

## Laravel

- https://laravel.com/docs/5.5/migrations
- use timestamp to order migration (if I remember it correctly)
- migration implementation in [DatabaseMigrationRepository](https://github.com/laravel/framework/blob/5.6/src/Illuminate/Database/Migrations/DatabaseMigrationRepository.php)
  - `compileTableExists` for different database in [Schema/Grammars](https://github.com/laravel/framework/tree/5.6/src/Illuminate/Database/Schema/Grammars)
  - `createRepository` creates the migration table
  
````php
class DatabaseMigrationRepository implements MigrationRepositoryInterface
{
    /**
     * Create the migration repository data store.
     *
     * @return void
     */
    public function createRepository()
    {
        $schema = $this->getConnection()->getSchemaBuilder();

        $schema->create($this->table, function ($table) {
            // The migrations table is responsible for keeping track of which of the
            // migrations have actually run for the application. We'll create the
            // table to hold the migration file's path as well as the batch ID.
            $table->increments('id');
            $table->string('migration');
            $table->integer('batch');
        });
    }
    
    /**
     * Get the completed migrations with their batch numbers.
     *
     * @return array
     */
    public function getMigrationBatches()
    {
        return $this->table()
                ->orderBy('batch', 'asc')
                ->orderBy('migration', 'asc')
                ->pluck('batch', 'migration')->all();
    }
        
    /**
     * Log that a migration was run.
     *
     * @param  string  $file
     * @param  int  $batch
     * @return void
     */
    public function log($file, $batch)
    {
        $record = ['migration' => $file, 'batch' => $batch];

        $this->table()->insert($record);
    }
    
    /**
     * Remove a migration from the log.
     *
     * @param  object  $migration
     * @return void
     */
    public function delete($migration)
    {
        $this->table()->where('migration', $migration->migration)->delete();
    }

}
````
  
````php
<?php
class CreateFlightsTable extends Migration
{
    public function up()
    {
        Schema::create('flights', function (Blueprint $table) {
            $table->increments('id');
            $table->string('name');
            $table->string('airline');
            $table->timestamps();
        });
    }

    public function down()
    {
        Schema::drop('flights');
    }
}
````

````php
<?php
class DatabaseSeeder extends Seeder
{
    public function run()
    {
        DB::table('users')->insert([
            'name' => str_random(10),
            'email' => str_random(10).'@gmail.com',
            'password' => bcrypt('secret'),
        ]);
    }
}
````

## Rails

- http://edgeguides.rubyonrails.org/active_record_migrations.html
- http://edgeguides.rubyonrails.org/active_record_migrations.html#migrations-and-seed-data

````ruby
class CreateProducts < ActiveRecord::Migration[5.0]
  def change
    create_table :products do |t|
      t.string :name
      t.text :description
      t.timestamps
    end
  end
end
````