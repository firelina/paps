Метод: POST

URL: /api/v1/users

Описание: Создание нового пользователя.

Тело запроса:

Пример:

`{
    "username": "Имя", - строка
    "password": "Пароль" - строка
}`

Ответ:

`{
    "user_id": 1
}`

Статусы ответов:

Код 201: Пользователь успешно создан.

Код 400: Неверные данные.

Метод: GET

URL: /api/v1/users/{id}

Описание: Получение пользователя.

Ответ:
`{
    "id": 1,
    "username": "Имя"
}`

Статусы ответов:

Код 200: Успешный запрос.

Код 404: Пользователь не найден.

Метод: GET

URL: /api/v1/scenarios

Описание: Получение списка всех учебных сценариев.

Ответ:
`[
    {
        "id": 1,
        "title": "Название сценария",
        "description": "Описание сценария"
    }
]`

Статусы ответов:

Код 200: Успешный запрос.

Код 500: Внутренняя ошибка сервера.

Метод: POST

URL: /api/v1/scenarios

Описание: Создание нового учебного сценария.

Тело запроса:
`{
    "title": "Название сценария", - строка
    "description": "Описание сценария" - строка
}`

Ответ:
`{
    "scenario_id": 1
}`

Статусы ответов:

Код 201: Сценарий успешно создан.

Код 400: Неверные данные.

Метод: PUT

URL: /api/v1/scenarios/{id}

Описание: Обновление существующего учебного сценария.

Тело запроса:
`{
    "title": "Новое название",
    "description": "Новое описание" 
}`

Ответ:
`{
    "id": 1,
    "title": "Новое название",
    "description": "Новое описание"
}`

Статусы ответов:

Код 200: Сценарий успешно обновлен.

Код 404: Сценарий не найден.

Метод: DELETE

URL: /api/v1/scenarios/{id}

Описание: Удаление учебного сценария.

Статусы ответов:

Код 204: Сценарий успешно удален.

Код 404: Сценарий не найден.

Метод: POST

URL: /api/v1/materials

Описание: Загрузка материалов (текстовых или аудиофайлов).

Тело запроса: 
`{
    "name": "Название материала" - строка
}`

Ответ:
`{
    "material_id": 1
}`

Статусы ответов:

Код 201: Материал успешно загружен.

Код 400: Неверные данные.

Метод: GET

URL: /api/v1/materials

Описание: Получение списка всех загруженных материалов.

Ответ:
`[
    {
        "id": 1,
        "name": "Название материала"
    }
]`

Статусы ответов:

Код 200: Успешный запрос.

Код 500: Внутренняя ошибка сервера.

Добавление пользователя

![post_user.PNG](post_user.PNG)

![post_user_tests.PNG](post_user_tests.PNG)

Получение пользователя

![get_user.PNG](get_user.PNG)
![get_user_tset.PNG](get_user_tset.PNG)

Добавление материала

![post_material.PNG](post_material.PNG)
![post_material_test.PNG](post_material_test.PNG)

Получение материала

![get_material.PNG](get_material.PNG)
![get_material_test.PNG](get_material_test.PNG)

Создание сценария

![post_scenario.PNG](post_scenario.PNG)
![post_scenario_test.PNG](post_scenario_test.PNG)

Получение сценария

![get_scen.PNG](get_scen.PNG)
![get_scen_test.PNG](get_scen_test.PNG)

Обновление сценария

![put_scen.PNG](put_scen.PNG)
![put_scen_test.PNG](put_scen_test.PNG)

Удаление сценария

![delete_scen.PNG](delete_scen.PNG)
![delete_scen_test.PNG](delete_scen_test.PNG)