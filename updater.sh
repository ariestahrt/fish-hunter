# get last modified date of file 'trigger_update'
last_modified=$(stat -c %y trigger_update)

# infinite loop
while true; do
    # get last modified date of file 'trigger_update'
    last_modified_new=$(stat -c %y trigger_update)
    # if last modified date of file 'trigger_update' has changed
    if [ "$last_modified" != "$last_modified_new" ]; then
        # trigger_update last modified date
        last_modified=$last_modified_new
        # terminate old process
        pkill dist

        # download packages
        go mod download

        # build binary
        go build -o dist

        # restart service
        ./dist 1>/dev/null 2>&1
    fi
    # sleep for 1 second
    sleep 1
done