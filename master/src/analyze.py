import numpy
import math
from scipy.special import dawsn as F
from scipy.special import fresnel

def fres(t):
	s, c = fresnel(math.sqrt(2.0 * t / math.pi))
	return math.cos(t) * c + math.sin(t) * s

#taken from paper
solns = {
	(0, 0): lambda t: math.exp(-1*t),
	(0, 1): lambda t: 3.0*t**2 + t**3,
	(0, 2): lambda t: math.sin(t) + math.cos(t),
	(1, 0): lambda t: 2.0*math.exp(-1*t) - 1.0,
	(1, 1): lambda t: 3.0*t**2 - (t**4 / 4.0),
	(1, 2): lambda t: 2.0*math.cos(t) - 1.0,
	(2, 0): lambda t: (1.0 / math.pi) * (math.sqrt(t) + (1.0 - 2.0 * t) * F(math.sqrt(t))),
	(2, 1): lambda t: (16.0 / (5.0 * math.pi)) * math.pow(t, 5.0 / 2.0),
	(2, 2): lambda t: math.sqrt(2.0 / math.pi) * fres(t),
}

def error(method, n, top, kernel, fn):
	errs = numpy.zeros(n)
	i = 0
	with open("results/%s.%d.%d.%d.%d.log" % (method, kernel, fn, n, top)) as f:
		for line in f:
			x, y = [float(z.strip()) for z in line.split("\t")]
			errs[i] = solns[(kernel, fn)](x) - y
			i += 1
	h = float(top) / float(n)
	return math.sqrt(h * sum(errs**2))

ns = [10, 100, 1000, 10000]
tops = [1, 2, 5]
methods = ["exact", "lubich", "mid", "rect"]
kernels = [0,1,2]
fns = [0,1,2]

for kernel in range(3):
	for fn in range(3):
		print "kernel%d\tf%d" % (kernel + 1, fn + 1)
		for method in ["exact", "lubich", "mid", "rect"]:
			print "\t%s:" % method
			print "\t\tStepsize\tTop\tError\t\t\tRatio"
			for top in tops:
				prev = 0
				for n in ns:
					err = error(method, n, top, kernel, fn)
					print "\t\t%f\t%d\t%.16f" % (float(top) / float(n), top, err),
					if prev != 0:
						print "\t%.16f" % math.log10(prev / err)
					else:
						print
					prev = err
				print

import matplotlib.pyplot as plt
def plot(method, kernel, fn, n, top, save=True):
	with open("results/%s.%d.%d.%d.%d.log" % (method, kernel, fn, n, top)) as f:
		if (method == "rect" or method == "lubich"):
			n -= 1

		x = numpy.zeros(n)
		y1 = numpy.zeros(n)
		y2 = numpy.zeros(n)
		err = numpy.zeros(n)

		if (method == "rect" or method == "lubich"):
			n += 1

		i = 0
		for line in f:
			xv, yv = [float(z.strip()) for z in line.split("\t")]
			sln = solns[(kernel,fn)](xv)
			x[i], y1[i], y2[i], err[i] = xv, yv, sln, sln - yv
			# print "%.8f\t%.8f\t%.8f\t%.8f" % (x[i], y1[i], y2[i], err[i])
			i += 1

		err = numpy.sqrt(err**2)

		plt.clf()
		plt.figure(1)
		plt.plot(x,y1)
		plt.plot(x,y2)
		plt.plot(x,err)
		if save:
			plt.savefig("figs/%s.%d.%d.%d.%d.png" % (method, kernel, fn, n, top))
		else:
			plt.show()

from itertools import product
for comb in product(methods, kernels, fns, ns, tops):
	plot(*comb)