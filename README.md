# CS411 Spring 2027 · Lab 6: The Partition Simulator

Starter template: https://github.com/cwillpinto/cs411-spring2027-lab06-the-partition-simulator

Read the included handout. Complete the CP/AP read policies in `policies.go` and clockwise node selection in `hashring.go`. The replica update and healing model is supplied in `simulator.go`; explain its version tie rule in your trace.

```sh
go test ./...
go test -race ./...
```

The unchanged starter compiles but fails public policy and hash-ring checks. Public tests are representative, not exhaustive.

## Work in your own GitHub repository

The instructor starter is an upstream source. Keep your completed work in a repository under your own GitHub account, then submit the repository URL and the evidence listed in the assignment handout through Blackboard.

If Blackboard links a template repository, choose **Use this template**, create the new repository under your account, and clone that new repository. Check `git remote -v` before pushing; `origin` must be your repository.

For a direct clone, first create an empty repository under your account. Replace the placeholders below with the starter link supplied in Blackboard and your own empty repository URL:

```sh
git clone https://github.com/cwillpinto/cs411-spring2027-lab06-the-partition-simulator
cd <cloned-directory>
git remote rename origin upstream
git remote add origin <your-empty-repository-url>
git remote -v
git push -u origin HEAD
```

Commit and push your completed work to your own `origin`. Do not push to the instructor's upstream repository. Use the exact submission files and verification commands in the assignment handout.
