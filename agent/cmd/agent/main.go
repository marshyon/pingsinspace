package main

import (
	"log"

	architecture "github.com/marshyon/pingsinspace/agent"

	systemexec "github.com/marshyon/pingsinspace/agent/systemExec"
)

func main() {
	log.Println("starting ...")
	sexec := systemexec.Job{}

	js := architecture.NewJobService(sexec)
	js.Run("test command to run", 1)
	log.Println("end run.")

}

// var workers = 100
// var summaryList string

// type JobList []struct {
// 	Org              string `json:"org"`
// 	OrgID            int    `json:"org_id"`
// 	Type             string `json:"type"`
// 	TypeID           int    `json:"type_id"`
// 	Name             string `json:"name"`
// 	NameID           int    `json:"name_id"`
// 	CmdCheckCommand  string `json:"cmd_check_command"`
// 	CmdServer        string `json:"cmd_server"`
// 	CmdURL           string `json:"cmd_url"`
// 	CmdStringToMatch string `json:"cmd_string_to_match"`
// 	CmdToRun         string `json:"cmd_to_run"`
// 	Cmd              string `json:"cmd"`
// }

// type jobRequest struct {
// 	orgID      int
// 	typeID     int
// 	nameID     int
// 	jobName    string
// 	jobCommand string
// }

// type jobRequestList []jobRequest

// type jobResult struct {
// 	orgID      int
// 	typeID     int
// 	nameID     int
// 	jobName    string `json: "job_name"`
// 	jobCommand string `json: "job_command"`
// 	jobOutput  string `json: "job_output"`
// 	jobStatus  int    `json: "job_status"`
// }

// func readJobsList(file string) (jl JobList) {
// 	jsonFile, err := os.Open(file)
// 	if err != nil {
// 		log.Fatalf("failed to open input file %s : %s\n", file, err)
// 	}
// 	defer jsonFile.Close()
// 	byteValue, err := ioutil.ReadAll(jsonFile)
// 	if err != nil {
// 		log.Fatalf("can't read input file [%s] : %s\n", file, err)
// 	}
// 	err = json.Unmarshal(byteValue, &jl)
// 	if err != nil {
// 		log.Fatalf("failed to unmarshall json : %s\n", err)
// 	}
// 	return jl
// }
// func createJobRequestList(jobsList JobList) (jobReqList jobRequestList) {
// 	for _, job := range jobsList {
// 		checkCommand := strings.Replace(job.Cmd, "{{cmd_check_command}}", job.CmdCheckCommand, -1)
// 		checkCommand = strings.Replace(checkCommand, "{{cmd_server}}", job.CmdServer, -1)
// 		checkCommand = strings.Replace(checkCommand, "{{cmd_url}}", job.CmdURL, -1)
// 		checkCommand = strings.Replace(checkCommand, "{{cmd_string_to_match}}", job.CmdStringToMatch, -1)
// 		jobReqList = append(jobReqList, jobRequest{
// 			orgID:      job.OrgID,
// 			typeID:     job.TypeID,
// 			nameID:     job.NameID,
// 			jobName:    job.Name,
// 			jobCommand: checkCommand,
// 		})
// 	}
// 	return jobReqList
// }

// func main() {

// 	// get jobs to run
// 	jobsList := readJobsList("jobs.json")
// 	jl := createJobRequestList(jobsList)

// 	// create channels
// 	jobsChannel := make(chan jobRequest)
// 	resultChannel := make(chan jobResult)

// 	// create workers (that will do work asynchronously)
// 	for i := 0; i < workers; i++ {
// 		go worker(jobsChannel, resultChannel, i)
// 	}

// 	// create generators (to send jobs to workers)
// 	for _, j := range jl {
// 		go generator(j, jobsChannel)

// 	}

// 	// receive back results from the workers and send them off to the time series database

// 	csvText := ""
// 	for i := 0; i < len(jl); i++ {
// 		res := <-resultChannel
// 		jobOutputTrimmed := strings.TrimSuffix(res.jobOutput, "\n")
// 		currentTime := time.Now()
// 		currTimeStr := fmt.Sprintf("%s+00", currentTime.Format("2006-01-02 15:04:05"))
// 		uuid := generateUUID()
// 		csvText += fmt.Sprintf("%s\t%d\t%s\t%d\t%d\t%s\t%s\t%d\n", uuid, res.orgID, jobOutputTrimmed, res.typeID, res.jobStatus, currTimeStr, currTimeStr, res.nameID)

// 	}

// 	// output results
// 	fmt.Printf("%s", csvText)
// }

// func generateUUID() string {
// 	out, err := exec.Command("uuidgen").Output()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	outStr := strings.TrimSuffix(string(out), "\n")
// 	return outStr
// }

// func runCheck(job string) (jobStatus int, jobOutput string, err error) {
// 	cmdArgs := strings.Split(job, " ")
// 	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
// 	cmdOutput := &bytes.Buffer{}
// 	cmdError := &bytes.Buffer{}
// 	cmd.Stdout = cmdOutput
// 	cmd.Stderr = cmdError
// 	var returnString string
// 	err = cmd.Start()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	done := make(chan error, 1)
// 	go func() {
// 		done <- cmd.Wait()
// 	}()

// 	select {
// 	case <-time.After(30 * time.Second):
// 		if err := cmd.Process.Kill(); err != nil {
// 			returnString := fmt.Sprintf("failed to kill [%s]", job)
// 			return 2, returnString, errors.New("failed to kill process")
// 		}
// 		returnString := fmt.Sprintf("Process timed out [%s]", job)
// 		return 2, returnString, errors.New("process killed as timeout reached")
// 	case err := <-done:
// 		if err != nil {
// 			combinedOutput := fmt.Sprintf("%s %s %s", cmdError.Bytes(), err, cmdOutput.Bytes())
// 			re := regexp.MustCompile("([0-9])")
// 			errStr := fmt.Sprintf("%s", err)
// 			strMatch := re.FindAllString(errStr, -1)
// 			i, err := strconv.Atoi(strMatch[0])
// 			if err != nil {
// 				i = 3
// 			}
// 			c := strings.TrimSpace(combinedOutput)

// 			return i, c, errors.New("command completed with errors")
// 		} else {
// 			returnString = string(cmdOutput.Bytes())
// 		}
// 	}

// 	return 0, returnString, nil
// }

// func worker(jobsChannel chan jobRequest, resultChannel chan jobResult, id int) {
// 	for {

// 		job := <-jobsChannel
// 		jobName := job.jobName
// 		jobCommand := job.jobCommand

// 		jobStatus, jobOutput, err := runCheck(jobCommand)
// 		res := new(jobResult)
// 		if err == nil {
// 			res.jobName = jobName
// 			res.jobCommand = jobCommand
// 			res.jobOutput = jobOutput
// 			res.jobStatus = jobStatus
// 			res.orgID = job.orgID
// 			res.typeID = job.typeID
// 			res.nameID = job.nameID
// 		} else {
// 			errorStr := fmt.Sprintf("%s", err)
// 			res.jobName = jobName
// 			res.jobCommand = jobCommand
// 			res.jobOutput = errorStr + " " + jobOutput
// 			res.jobStatus = jobStatus
// 		}

// 		resultChannel <- *res
// 	}
// }

// func generator(job jobRequest, jobsChannel chan jobRequest) {
// 	jobsChannel <- job
// }
