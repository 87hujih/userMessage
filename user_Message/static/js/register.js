//用户注册响应检测
document.addEventListener('DOMContentLoaded', function () {
    const usernameInput = document.getElementById('username');
    const passwordInput = document.getElementById('password');
    const usernameRemind = document.getElementById('usernameRemind');
    const passwordRemind = document.getElementById('passwordRemind');
    const eyeIcon = document.getElementById('eyeIcon');
    let eye_Status=true;
    // 点击输入框时显示提示信息
    usernameInput.addEventListener('focus', function () {
        usernameRemind.style.display = 'block';
    });

    passwordInput.addEventListener('focus', function () {
        passwordRemind.style.display = 'block';
    });

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
        if (eye_Status){
            passwordInput.type = 'text';
            eyeIcon.src = "../static/img/icon-eye-invisable@2x.png";
            eye_Status=false;
        }else {
            passwordInput.type = 'password';
            eyeIcon.src = "../static/img/icon-eye-visable@2x.png";
            eye_Status=true;
        }
    });

    // 点击其他地方或输入内容后隐藏提示信息
    usernameInput.addEventListener('blur', function () {
        if (usernameInput.value.trim() !== '') {
            usernameRemind.style.display = 'none';
        }
    });

    passwordInput.addEventListener('blur', function () {
        if (passwordInput.value.trim() !== '') {
            passwordRemind.style.display = 'none';
        }
    });
    
});

