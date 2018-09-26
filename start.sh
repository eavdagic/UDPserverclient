port="8011"
address="127.0.0.1:8080"

while getopts "p:a:" flag; do
	case "${flag}" in
		p) port="${OPTARG}" ;;
		a) address="${OPTARG}" ;;
	esac
done

./server/server -address=$address &
sleep 2 &
./client/client -port=$port -address=$address



