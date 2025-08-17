<template>
  <view class="login-container">
    <!-- 顶部紫色区域 -->
    <view class="header">
      <view class="icon-container">
        <image class="icon" src="/static/briefcase.png"></image>
      </view>
      <view class="welcome-text">欢迎回来</view>
      <view class="sub-text">请登录您的账号以继续</view>
    </view>
    
    <!-- 登录表单区域 -->
    <view class="form-container">
      <!-- 账号输入框 -->
      <view class="input-item">
        <text class="input-icon">👤</text>
        <input type="text" placeholder="请输入账号" v-model="username" class="input-field" />
      </view>
      
      <!-- 密码输入框 -->
      <view class="input-item">
        <text class="input-icon">🔒</text>
        <input type="password" placeholder="请输入密码" v-model="password" class="input-field" />
      </view>
      
      <!-- 记住我和忘记密码 -->
      <view class="options-row">
        <view class="remember-me">
          <checkbox :checked="rememberMe" @tap="toggleRememberMe" />
          <text>记住我</text>
        </view>
        <view class="forgot-password" @tap="forgotPassword">忘记密码?</view>
      </view>
      
      <!-- 登录按钮 -->
      <button class="login-btn" @tap="login">登录</button>
      
      <!-- 注册链接 -->
      <view class="register-link">
        <text>还没有账号? </text>
        <text class="register-text" @tap="goToRegister">立即注册</text>
      </view>
    </view>
  </view>
</template>

<script setup>
import { ref } from 'vue';

// 响应式数据
const username = ref('');
const password = ref('');
const rememberMe = ref(false);

// 切换记住我状态
const toggleRememberMe = () => {
  rememberMe.value = !rememberMe.value;
};

// 忘记密码
const forgotPassword = () => {
  uni.showToast({
    title: '忘记密码功能开发中',
    icon: 'none'
  });
};

// 登录方法
const login = () => {
  if (!username.value || !password.value) {
    uni.showToast({
      title: '请输入账号和密码',
      icon: 'none'
    });
    return;
  }
  
  // 这里添加登录逻辑
  uni.showLoading({
    title: '登录中...'
  });
  
  // 模拟登录请求
  setTimeout(() => {
    uni.hideLoading();
    
    const users = uni.getStorageSync('users') || [];
    const user = users.find(u => u.account === username.value && u.password === password.value);
    
    if (user) {
      uni.showToast({
        title: '登录成功',
        icon: 'success'
      });
      
      // 保存当前登录用户
      uni.setStorageSync('currentUser', user.account);
      
      // 登录成功后跳转到首页
      uni.switchTab({
        url: '/pages/index/index'
      });
    } else {
      uni.showToast({
        title: '账号或密码错误',
        icon: 'none'
      });
    }
  }, 1500);
};

// 跳转到注册页面
const goToRegister = () => {
  uni.navigateTo({
    url: '/pages/register/register'
  });
};
</script>

<style>
.login-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
}

.header {
  background-color: #6c5ce7;
  padding: 50rpx 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  color: white;
  border-bottom-left-radius: 30rpx;
  border-bottom-right-radius: 30rpx;
}

.icon-container {
  width: 120rpx;
  height: 120rpx;
  background-color: rgba(255, 255, 255, 0.3);
  border-radius: 30rpx;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 20rpx;
}

.icon {
  width: 60rpx;
  height: 60rpx;
}

.welcome-text {
  font-size: 40rpx;
  font-weight: bold;
  margin-bottom: 10rpx;
}

.sub-text {
  font-size: 28rpx;
  opacity: 0.8;
}

.form-container {
  padding: 40rpx;
  background-color: #fff;
  border-radius: 30rpx;
  margin-top: -20rpx;
  flex: 1;
}

.input-item {
  display: flex;
  align-items: center;
  border: 1px solid #e0e0e0;
  border-radius: 50rpx;
  padding: 20rpx 30rpx;
  margin-bottom: 30rpx;
}

.input-icon {
  margin-right: 20rpx;
  color: #999;
}

.input-field {
  flex: 1;
  font-size: 28rpx;
}

.options-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 40rpx;
  font-size: 26rpx;
}

.remember-me {
  display: flex;
  align-items: center;
}

.forgot-password {
  color: #6c5ce7;
}

.login-btn {
  background-color: #6c5ce7;
  color: white;
  border-radius: 50rpx;
  padding: 20rpx 0;
  font-size: 32rpx;
  margin-bottom: 40rpx;
}

.register-link {
  text-align: center;
  font-size: 26rpx;
  color: #666;
}

.register-text {
  color: #6c5ce7;
}
</style>