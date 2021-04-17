void main() {
  List team = ['valdir', 'andre', 'mateus'];

  var nome = 'Cristiano';
  print('Bem Vindo ' + nome + '\n\nO resto do time é:\n');

  for (var individuo in team) {
    print(' - ' + individuo);
  }
}
