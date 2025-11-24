# API Endpoints
## Посты
Не требуют авторизацию:
- `GET /post/:id`: получить пост по id.
- `GET /post?{getPostsParameters...}`: получить посты по параметрам getPostsParameters.

Требуют авторизацию:
- `POST /post`: создать пост.
- `PUT /post/:id`: изменить данные о посте по id.
- `DELETE /post/:id`: удалить пост по id.

### GetPostsParameters
- `ids: []int`
- `page: int`
- `perPage: int`
- `sortField: int` (разрешены: "created_at", "updated_at")
- `sortOrder: int` (разрешены: "asc", "desc")
- `name: int` (ищутся посты с именем, в котором содержится это поле)
- `authorId: int`

## Пользователи
Не требуют авторизацию:
- `GET /user/:id`: получить информацию о пользователе по id.
- `POST /user`: создать пользователя.

Требуют авторизацию:
- `PUT /user/:id`: изменить данные о пользователе по id.
- `DELETE /user/:id`: удалить пользователя по id.

## Аутентификация:
- `POST /auth/signin`: регистрация.
- `POST /auth/login`: вход.
- `POST /auth/refresh`: обновление токенов по refresh-токену (требует refresh-токен).
- `POST /auth/logout`: выход из аккаунта.
