import csv
import matplotlib.pyplot as plt

vars = []
seq = []
par = []

with open("results.csv") as f:
    reader = csv.DictReader(f)
    for row in reader:
        vars.append(int(row["vars"]))
        seq.append(float(row["seq_ms"]))
        par.append(float(row["par_ms"]))

plt.plot(vars, seq, label="Séquentiel")
plt.plot(vars, par, label="Parallèle")
plt.xlabel("Nombre de variables")
plt.ylabel("Temps (ms)")
plt.legend()
plt.grid()
plt.show()
