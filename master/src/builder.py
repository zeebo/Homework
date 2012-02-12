import subprocess

for points in [10, 100, 1000, 10000]:
	for top in [1, 2, 5]:
		print "Points: %d\tTop: %d" % (points, top)
		subprocess.call(["/bin/bash", "runner.sh"], env={
			"N": str(points),
			"TOP": str(top),
			"PATH": "/Users/zeebo/Code/go/bin",
		})