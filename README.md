# GoRecordurbate WebUI
This project offers a simple, self-hosted web UI for managing and recording streams.

__API NOTE__: The api is basically un-tested after adding login credentials. The plan is to use API key in the future for credentials via API. This will probably not be done before the initial release.
## Core Features:
- Login with cookies for each client (prevents unsupervised access)
- Start, restart, and stop recordings
- Add/delete streamers, with import/export options
- View logs and recorded videos directly in the web UI
- Docker and service examples for easier setup
- Manage and add multiple users for webUI login
- Add new streamers without restarting the recorder
- View active recorders
  
## Usage

|Username|Password|
|-|-|
|`admin`| `password`|

### Setup
- Copy [`.env.example`](https://github.com/luna-nightbyte/Recordurbate-WebUI/blob/main/.env.example) to `.env` and add your own session key. 
    - Can be generated with: `head -c 32 /dev/urandom | base64`
    - Example:
      ```bash
      user@user:~/Recordurbate$ head -c 32 /dev/urandom | base64
      # Output:
      Fl60B6sTqAUARyDiC6GIor8AIu6QXLF2RMvWK1Wz3eE=
      ```

#### Optional config settings
The main settings can be found in [`settings.json`](https://github.com/luna-nightbyte/Recordurbate-WebUI/blob/main/internal/settings/settings.json):
```json
{
  "app": {
    "port": 8055,
    "loop_interval_in_minutes": 2,
    "video_output_folder": "output/videos",
    "rate_limit": {
      "enable": true,
      "time": 5
    },
    "default_export_location": "./output/list.txt"

  },
  "youtube-dl": {
    "binary": "youtube-dl"
  },
  "auto_reload_config": true
}
```
#### Reset password
To change forgotten password, start the program with the `reset-pwd` argument. I.e:
```
./GoRecordurbate reset-pwd admin newpassword 
```
New login for the user `admin` would then be `newpassword`
### Build
Building the code wil create a binary for your os system. Golang is [cross-compatible](https://go.dev/wiki/GccgoCrossCompilation) for windows, linux and mac.
```bash
go mod init GoRecordurbate # Only run this line once
go mod tidy
go build
./GoRecordurbate #windows will have 'GoRecordurbate.exe'
```
### Source
```bash
go mod init GoRecordurbate # Only run this line once
go mod tidy
go run main.go
```

## Notes
This is un-tested on Windows and Mac, but golang is cross-compatible which means that this should run just as fine on Windows as on Linux.

## WebUI (Will probably be modified)


<p align="center">
  <img src="https://github.com/user-attachments/assets/35e4633b-702b-45f9-9075-a8522a6b334b" alt="Login page"/>
  <img src="https://github.com/user-attachments/assets/cc3f013f-a530-4629-99d4-02b72a54599d" alt="Streamers tab"/>
  <img src="https://github.com/user-attachments/assets/86b72aba-e3a8-424a-b094-8e23b66a233b" alt="Logs tab"/>
  <img src="https://github.com/user-attachments/assets/7038689f-0c67-4ede-87aa-1d5a63a720c2" alt="Recorders tab"/>
  <img src="https://github.com/user-attachments/assets/30bc2424-e3c6-49f7-8f97-b4025698e234" alt="Settings tab"/>
</p>

## Other

### Todo, but not planned/prioritized 
- Log online-time of streamers and save to csv for graph plotting. Can help understand the work-hours of different streamers.
- Option to login to the streaming site and use follower list instead of config?
- Option for max video length (and size?)
- ~~headless mode without webui~~ (Abandoned)
- Move frontend to Vue (?)
- Build a default docker image
- Individual recorders in UI
  - Stop/Restart individual recorders
  - view current recording length
  - view current recording video





## Thanks

Special thanks to [oliverjrose99](https://github.com/oliverjrose99) for the inspiration and their work on [Recordurbate](https://github.com/oliverjrose99/Recordurbate)
