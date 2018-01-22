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