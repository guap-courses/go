# Домашнее задание №1

Необходимо реализовать функцию, которая будет обрабатывать слайс целых чисел с возможностью настройки различных параметров. Функция будет выполнять следующие операции над входным слайсом:

1. Удаление дубликатов.
2. Сортировка чисел.
3. Фильтрация чисел по минимальному значению.

### Требования к реализации

1. Функция должна принимать слайс целых чисел и следующие настройки:
   - сортировка: необходимо ли сортировать числа (по умолчанию false).
   - удаление дубликатов: нужно ли удалять дубликаты (по умолчанию true).
   - минимальное значение: минимальное значение для фильтрации, числа меньще этого значения должны отбасываться (по умолчанию 0).

2. Функция должна возвращать результирующий слайс целых чисел.

3. Сигнатура функции не должна изменяться при добавлении новых параметров.

  Т.е. вот так плохо:
  ```go
  func processNumbers(sort true, removeDuplicates true, minValue int, ...)
  ```
  Стоит придумать решение, которое будет работать как-то так (**это псевдокод**):
  ```go
  processNumbers(nums)    // параметры для работы функции будут использоваться по умолчанию
  processNumbers(nums disableSorting)  // указан только один параметр, отвечающий за сортировку
  processNumbers(nums sort=true minvalue=20)  // указано два параметра
  ...
  ```

- Можно использовать только типы данных, которые прошли на этом занятии. Так что никаких структур и интерфейсов. В ваших руках все скалярные типы, а также массивы, словари и функции.
- Можно использовать только стандартную библиотеку Go. Учтите, что это задание можно сделать, используя только "fmt"