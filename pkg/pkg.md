Код библиотек, пригодных для использования в сторонних приложениях. 
(например, /pkg/mypubliclib). Другие проекты будут импортировать эти библиотеки, ожидая их 
автономной работы, поэтому стоит подумать дважды, прежде чем класть сюда какой-нибудь код :-)
Заметьте, что использование директории internal - более оптимальный способ не дать импортировать
внутренние пакеты, потому что это обеспечит сам Golang. Директория /pkg - всё еще хороший путь
дать понять, что код в этой директории могут безопасно использовать другие. Пост I'll take pkg
over internal в блоге Трэвиса Джеффери предоставляет хороший обзор директорий pkg и internal и 
когда есть смысл их использовать.

Существует возможность группировать код на Golang в одном месте, когда ваша корневая директория
содержит множество не относящихся к Go компонентов и директорий, что позволит облегчить работу
с разными инструментами Go.

Ознакомьтесь с директорией /pkg, если хотите увидеть, какие популярные репозитории используют 
такой шаблон организации проекта. Несмотря на его распространенность, он не был принят всеми, 
а некоторые в обществе Go и вовсе не рекомендует его использовать.

Вы можете не использовать эту директорию, если проект совсем небольшой и добавление нового 
уровня вложенности не имеет практического смысла (разве что вы сами этого не хотите :-)).
Подумайте об этом, когда ваша корневая директория начинает слишком сильно разрастаться, 
особенно, если у вас много компонентов, написанных не на Go.