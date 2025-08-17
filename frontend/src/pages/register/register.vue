<script setup>
import { ref } from 'vue';

// 响应式数据
const formData = {
  account: ref(''),
  password: ref(''),
  confirmPassword: ref('')
};

// 注册方法
const handleRegister = () => {
  // 表单验证
  if (!formData.account.value || !formData.password.value || !formData.confirmPassword.value) {
    uni.showToast({
      title: '请填写完整信息',
      icon: 'none'
    });
    return;
  }

  if (formData.password.value !== formData.confirmPassword.value) {
    uni.showToast({
      title: '两次密码输入不一致',
      icon: 'none'
    });
    return;
  }
  
  // 模拟注册请求
  uni.showLoading({
    title: '注册中...'
  });
  
  setTimeout(() => {
    uni.hideLoading();
    
    // 将用户信息存储到本地
    const users = uni.getStorageSync('users') || [];
    const existingUser = users.find(user => user.account === formData.account.value);
    
    if (existingUser) {
      uni.showToast({
        title: '该账号已被注册',
        icon: 'none'
      });
      return;
    }
    
    users.push({
      account: formData.account.value,
      password: formData.password.value 
    });
    
    uni.setStorageSync('users', users);
    
    uni.showToast({
      title: '注册成功',
      icon: 'success'
    });
    
    // 注册成功后返回登录页
    setTimeout(() => {
      uni.navigateBack();
    }, 1500);
  }, 1500);
};

// 返回登录页
const handleBack = () => {
  uni.navigateBack();
};
</script>

<template>
  <view class="register-container">
    <!-- 顶部紫色区域 -->
    <view class="header">
      <view class="icon-container">
        <image class="icon" src="/static/briefcase.png"></image>
      </view>
      <view class="welcome-text">创建账号</view>
      <view class="sub-text">请填写以下信息完成注册</view>
    </view>

    <!-- 表单区域 -->
    <view class="form-container">
      <!-- 账号输入框 -->
      <view class="input-item">
        <text class="input-icon">👤</text>
        <input type="text" placeholder="请输入账号" v-model="formData.account.value" class="input-field" />
      </view>
      
      <!-- 密码输入框 -->
      <view class="input-item">
        <text class="input-icon">🔒</text>
        <input type="password" placeholder="请输入密码" v-model="formData.password.value" class="input-field" />
      </view>
      
      <!-- 确认密码输入框 -->
      <view class="input-item">
        <text class="input-icon">🔒</text>
        <input type="password" placeholder="请确认密码" v-model="formData.confirmPassword.value" class="input-field" />
      </view>
      
      <!-- 注册按钮 -->
      <button class="register-btn" @tap="handleRegister">注册</button>
      
      <!-- 返回登录 -->
      <view class="login-link">
        <text>已有账号? </text>
        <text class="login-text" @tap="handleBack">返回登录</text>
      </view>
    </view>
  </view>
</template>

<style>
.register-container {
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

.register-btn {
  background-color: #6c5ce7;
  color: white;
  border-radius: 50rpx;
  padding: 20rpx 0;
  font-size: 32rpx;
  margin-bottom: 40rpx;
}

.login-link {
  text-align: center;
  font-size: 26rpx;
  color: #666;
}

.login-text {
  color: #6c5ce7;
}
</style>
