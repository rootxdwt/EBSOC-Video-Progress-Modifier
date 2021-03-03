# EBSOC-Video-Progress-Modifier

This program can modify your EBS onlineclass video progress to whatever number you want,
using EBS's vulnerability which has the cryptokey and iv hardcoded in their source code.

EBS 온라인클래스의 영상 진행상태를 원하는 숫자로 설정할 수 있게 해주는 프로그램입니다, 
암호화 키와 iv가 소스에 그대로 하드코드되어있는 EBS의 취약점을 이용합니다.

# Instruction

1. go to the target video in ebs online class.
2. right-click the video and click inspect.
3. in the developer tools, scroll down and find the <script></script> tag (without any src attribute) and click it.
4. click the console tab and paste the content of ```paste.js``` .
5. copy the output and paste it in ```ebs.exe``` .
6. the program will close itself when everything is done.
