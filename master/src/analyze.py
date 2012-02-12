import numpy
import math

#taken from paper
solns = {
	(0, 0): lambda t: math.exp(-1*t),
	(0, 1): lambda t: 3.0*t**2 + t**3,
	(0, 2): lambda t: math.sin(t) + math.cos(t),
	(1, 0): lambda t: 2.0*math.exp(-1*t) - 1.0,
	(1, 1): lambda t: 3.0*t**2 - (t**4 / 4.0),
	(1, 2): lambda t: 2.0*math.cos(t) - 1.0,
}

def error(method, n, top, kernel, fn):
	errs = numpy.zeros(n)
	i = 0
	with open("results/%s.%d.%d.%d.%d.log" % (method, kernel, fn, n, top)) as f:
		for line in f:
			x, y = [float(z.strip()) for z in line.split("\t")]
			errs[i] = solns[(kernel, fn)](x) - y
			i += 1
	return math.sqrt(sum(errs**2))

ns, tops = [10, 100, 1000, 10000], [1, 2, 5]
for kernel in range(2):
	for fn in range(3):
		print "kernel%d\tf%d" % (kernel + 1, fn + 1)
		for method in ["exact", "lubich", "mid", "rect"]:
			print "\t%s:" % method
			print "\t\tStepsize\tTop\tError\t\tRatio"
			for top in tops:
				prev = 0
				for n in ns:
					err = error(method, n, top, kernel, fn)
					print "\t\t%f\t%d\t%f" % (float(top) / float(n), top, err),
					if prev != 0:
						print "\t%f" % (prev / err)
					else:
						print
					prev = err
				print
