notas = []

while True:
    x = input("Digite uma nota ou digite exit para terminar: ")

    if x == "exit":
        break

    notas.append(float(x))

print("A media é %.2f" % (sum(notas)/len(notas)))
