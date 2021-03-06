# STRESS TESTING
A test project for load testing.

## HOW TO WORK WITH APPLICATION
````
make app_up
make app_down
````

## RUN STRESS TESTS
````
make k6_test_jaeger
make k6_test_stress_spike
````

## VIEW JAEGER INFORMATION
Сервис jaeger доступен здесь: http://localhost:16686/

## VIEW RESULT OF STRESS TESTING IN GRAFANA
Результаты нагрузочного тестирования можно посмотреть в grafana: http://localhost:3000/

## ABOUT RESULTS
Процентиль — статистическая мера, указывающая значение, которое заданная случайная величина не превышает с указанной вероятностью. 
Например, фраза 95-й процентиль равен 7 означает, что 95% всех измеренных величин не достигает значения 7 
и только 5% всех измеренных величин превышает это значение.

## TEST TYPES
- Smoke testing.  
Это нагрузочный тест с минимальной нагрузкой. Например, чтобы проверить, что система вообще работает.
- Load testing.  
Создание некоторой нагрузки для системы и проверка, что все работает с адекватным количеством ошибок.
Цель - замерить производительность системы.
- Stress testing.  
Похоже на нагрузочное тестирование, но здесь цель - проверить доступность системы на пиковых нагрузках.
- Soak/Endurance testing.  
Тестирование системы под нагрузкой длительное время. Позволяет выявить проблемы с базой данных, 
утечки памяти и другие проблемы. 

