## Паттерн Фабричный метод

### Общие сведения

Порождающий паттерн. Полезен, если стоит задача избежать оператора ** new ** и система должна оставаться расширяемой, поэтому этот паттерн иногда называют обобщенным конструктором.

### Достоинства

- Создает объекты разных типов, позволяя системе оставаться независимо   как от самого процесса создания, так и от создаваемых объектов.

### Недостатки

- В случае введения нового типа объекта нужно создавать соответствующую фабрику.