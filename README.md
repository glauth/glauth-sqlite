# GLAuth Plugin

This is a GLAuth plugin; that is, a backend that are not compiled in GLAuth by default.

To quote 'Butonic' (Jörn Friedrich Dreyer):

> Just keep the 'lightweight' in mind.

To build either back-end, type
```
make plugin_name
```
where 'name' is the plugin's name; so, for instance: `make plugin_sqlite`

To build back-ends for specific architectures, specify `PLUGIN_OS` and `PLUGIN_ARCH` --
 For instance, to build the sqlite plugin for the new Mac M1s:
 ```
make plugin_sqlite PLUGIN_OS=darwin PLUGIN_ARCH=arm64
 ```

## Database Plugins

To use a database plugin, edit the configuration file (see pkg/plugins/sample-database.cfg) so that:

```
...
[backend]
  datastore = "plugin"
  plugin = "dynamic library you created using the previous 'make' command"
  database = "database connection string"
...
```
so, let's say you built the 'sqlite' plugin, you would now specify its library: `database = sqlite.so`

### SQLite, MySQL, Postgres

Tables:
- users, ldapgroups are self-explanatory
- includegroups store the 'includegroups' relationships
- othergroups, on the other hand, are a comma-separated list found in the users table (performance)

For documentation on seeding your database schema, and more, please visit this page:

https://glauth.github.io/docs/databases.html
