# Laravel-Sail-builder

- This is an alternative version of https://laravel.build.
- It helps you initialize a new Laravel Project with the desired version using Sail. (By default, Laravel.build will initialize the latest version.)
- Example:
  - http://localhost:8080/learn-laravel?version=10.0&with=pgsql,redis

```bash
curl -s "http://localhost:8080/learn-laravel?version=10.0&with=redis,mysql" | bash
```