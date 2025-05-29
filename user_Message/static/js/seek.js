document.addEventListener('DOMContentLoaded', function () {
    const passwordRemind = document.getElementById('seekPsdRemind');
    const passwordInput = document.getElementById('password');
    const eyeIcon = document.getElementById('eyeIcon');
    let eye_Status = true;
    passwordInput.addEventListener('focus', function () {
        passwordRemind.style.display = 'block';
    });
    passwordInput.addEventListener('blur', function () {
        if (passwordInput.value.trim() !== '') {
            passwordRemind.style.display = 'none';
        }
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

//用户名和密码格式验证
    const loginForm = document.getElementById("seekPsdForm");
    loginForm.addEventListener("submit", function (e) {
        const password = document.getElementById("password").value;
        const phone = document.getElementById("phone").value;
        if (!/^(?=.*[A-Za-z])(?=.*\d).{8,20}$/.test(password)) {
            alert("密码需8-20位，包含字母和数字");
            e.preventDefault();
        }
        if (!/^1[0-9]{10}$/.test(phone)) {
            alert("手机号格式不正确");
            e.preventDefault();
        }
    });
});