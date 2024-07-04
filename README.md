# Resource Databases Models Connector


## Sub-module
```azure

Mysql Databases 연결을 git submodule 방식으로 연결 사용  

```

### 폴더 설명
- /rdb_config: 설정파일 & Define File
- /rdb_helper: Helper, 
- /rdb_log: Log Database 연결
- /rdb_master: Master Database 연결

### 사용방법
1. DB 연결(Connect)

- Master Databases
```azure

ConnectMasterDatabases(masterHost *string, masterPort *int, maserUser *string, masterPass *string)

```
- Log Databases
```azure
ConnectLogDatabases(logHost *string, logPort *int, logUser *string, logPass *string)

```
***

2. Gorm으로 데이터 가져오기
- DB가 먼저 연결되어 있어야 합니다.(1번 연결후 사용)
- 데이터 모델을 이용한 데이터 가져오기
```azure
masterModel := res_model.MasterModel{
        Gorm: res.RdbConnMaster.DbConn,
}
rtRow := masterModel.GetMasterDeviceInfo(endpointId)
if rtRow == nil {
    return res_model.MasterModel{}, errors.New("no-data")
}

return rtRow, nil

```

***


********

#GIT 서브모듈
https://data-engineer-tech.tistory.com/20

********   

## Git Submodule 이란
```vertica
프로젝트를 할 때 코드를 깃에 올려서 관리를 하는데, 나중에 다른 곳에서 중복 코드를 만들지 않고 사용하려면 깃을 분리해서 코드를 나누는 것이 좋다.   
이럴 때를 위해서 git에서 제공해주는 것이 서브모듈(submodule)이다.   

출처: https://data-engineer-tech.tistory.com/20 [데이터 엔지니어 기술 블로그:티스토리]
```
   

## 기본 사용법
```vertica

서브모듈을 이용하면 프로젝트 서브폴더 생성과
# 서브모듈 정보를 기반으로 로컬 환경설정 파일을 만들어준다.
git submodule init

# 서브모듈의 리모트 저장소에서 데이터를 가져오고 Checkout을 한다.
git submodule update

```
