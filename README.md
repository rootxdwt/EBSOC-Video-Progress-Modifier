# EBSOC-Video-Progress-Modifier

This program can modify your EBS onlineclass video progress to whatever number you want,
using EBS's vulnerability which has the cryptokey and iv hardcoded in their source code.

EBS 온라인클래스의 영상 진행상태를 원하는 숫자로 설정할 수 있게 해주는 프로그램입니다, 
암호화 키와 iv가 소스에 그대로 하드코드되어있는 EBS의 취약점을 이용합니다.

# Notice

> EN

- this program is built for windows 10. if you are not using windows 10, download the source( `ebs.go` ) and compile it in your OS.
- this program is made only for educational purpose. we are not responsible for any malicious usage.

> KR

- 이 프로그램은 윈도우 10 전용으로 제작되었습니다. 다른 운영체제를 사용중이사라면 소스( `ebs.go` )를 다운받으시고 다시 컴파일해 사용해 주세요.
- 이 프로그램은 교육적인 용도로 제작되었습니다.

# Instruction

> EN

0. Download `ebs.exe` from the [latest release](https://github.com/rootxdwt/EBSOC-Video-Progress-Modifier/releases/).
1. go to the target video in ebs online class
2. right-click the video and click `Inspect`
3. in the developer tool, scroll down and find the `<script></script>` tag (without any src attribute) and click it
4. click the `Console` tab and paste the content of `paste.js` 
5. copy the output and paste it in `ebs.exe` 
6. the program will close itself when everything is done

> KR

0. [최신 릴리즈](https://github.com/rootxdwt/EBSOC-Video-Progress-Modifier/releases/)에서 `ebs.exe`를 다운받아주세요.
1. 조작하고 싶은 영상으로 이동해주세요.
2. 영상을 우클릭하시고 `검사` 를 클릭하세요.
3. `검사` 를 눌렀을때 나오는 개발자도구에서 스크롤을 내려 `<script></script>` 태그(src 속성이 없는) 를 찾고 클릭하세요.
4. 하단 `Console` 탭을 클릭하시고 `paste.js`의 내용을 붙여넣으세요.
5. 결과를 복사하시고 다운받으신 `ebs.exe`에 붙여넣으세요.
6. 모든 것이 끝나면, 프로그램이 자동으로 종료됩니다.
