дальнейшее развитие game-0
добавляем взаимодействие нескольких игроков
раньше был 1 игрок и игра была по сути синхронна
теперь добавляется второй ( и третий, четвёртый ) и они оба могут "вводить" команды
значит у нас появляется асинхронное взаимодействие
надо использовать каналы и горутины для обработки команд игроков в каждой комнате
Смотрите примеры с мультипрексорами для организации ввода

Большая часть команд так же относится к одному игроку, но появляется несколько новых команд:
"сказать" - транслируется всем другим игрокам в этой комнате
"сказать_игроку" - отправляет персональное сообщение другому игроку, остальные его не слышат
Если выполнить команду "осмотреться", что можно увидеть другого игрока
Это значит, что у комнаты появляется список игроков в ней, которым надо передать сообщение

Чтобы как-то отличать игроков друг от друга, теперь при инициализации надо задать имя ( функция NewPlayer )
Естественно, количество игроков не должно быть захардкожено и может меняться

У игрока появляется метод HandleInput, который принимает от него входящее сообщение. 
А так же метод HandleOutput, в который мы пишем то что игро должен видеть у себя на экране.
И метод GetOutput, который возвращает канал, принимающий сообщения
