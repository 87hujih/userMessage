document.addEventListener('DOMContentLoaded', function () {
    const passwordInput = document.getElementById('password');
    const eyeIcon = document.getElementById('eyeIcon');
    let eye_Status = true;
// 监听密码输入框的变化
    passwordInput.addEventListener('input', function () {
        // 如果密码不为空则显示眼睛图标
        if (passwordInput.value.length > 0) {
            eyeIcon.style.display = 'block';
        } else {
            eyeIcon.style.display = 'none';
            passwordInput.type = 'password';
            eyeIcon.src = "../static/img/icon-eye-visable@2x.png";
        }
    });

// 点击眼睛图标显示密码
    eyeIcon.addEventListener('click', function () {
        if (eye_Status) {
            passwordInput.type = 'text';
            eyeIcon.src = "../static/img/icon-eye-invisable@2x.png";
            eye_Status = false;
        } else {
            passwordInput.type = 'password';
            eyeIcon.src = "../static/img/icon-eye-visable@2x.png";
            eye_Status = true;
        }
    });
    const seekPsd=document.getElementById("forgotPsd")
    seekPsd.addEventListener("click",function (){

    })
});