# sem usar strings

dividendo = int(input("Digite um inteiro para dividir: "))
divisor = int(input("Digite um inteiro divisor: "))

print("\nParte inteira: %i\nParte decimal: %f\n" %
      (dividendo//divisor, (dividendo % divisor)/divisor))
