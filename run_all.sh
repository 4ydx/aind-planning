run=$1
if [[ "$run" -eq "" ]]; then
	echo "needs an argument"
	exit
fi
for i in 1 2 3 4 5 6 7 8 9 10; do
	echo "-- python run_search.py -p 1 -s $i"
	echo "-- python run_search.py -p 1 -s $i" >> log/$run-problem-1.log
	timeout 10m python run_search.py -p 1 -s $i >> log/$run-problem-1.log
done
for i in 1 2 3 4 5 6 7 8 9 10; do
	echo "-- python run_search.py -p 2 -s $i"
	echo "-- python run_search.py -p 2 -s $i" >> log/$run-problem-2.log
	timeout 10m python run_search.py -p 2 -s $i >> log/$run-problem-2.log
done
for i in 1 2 3 4 5 6 7 8 9 10; do
	echo "-- python run_search.py -p 3 -s $i"
	echo "-- python run_search.py -p 3 -s $i" >> log/$run-problem-3.log
	timeout 10m python run_search.py -p 3 -s $i >> log/$run-problem-3.log
done
