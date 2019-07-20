class Person():
  def __init__(self, name):
    self.name = name

  def talk(self):
    print('My name is %s' % self.name)

class Student(Person):
  def __init__(self, name, college):
    self.college = college
    super().__init__(name)

  def talk(self):
    super().talk()
    print('I go to %s' % self.college)

Catherine = Person('Catherine')
Hollis = Student('Hollis', 'Princeton')

Catherine.talk()
Hollis.talk()