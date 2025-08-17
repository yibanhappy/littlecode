<template>
  <view class="dashboard">
    <!-- 头部 -->
    <view class="header">
      <view class="logo">
        <image src="/static/1.png" class="logo-img"></image>
        <text class="title">个人工作台</text>
      </view>
      <view class="user-info">
        <view class="avatar">U</view>
        <text class="username">{{ username }}</text>
      </view>
    </view>

    <!-- 功能模块 -->
    <view class="modules">
      <!-- 备忘录模块 -->
      <view class="module memo-module">
        <view class="module-header">
          <text class="module-icon">📝</text>
          <text class="module-title">备忘录</text>
        </view>
        <scroll-view class="memo-list" scroll-y>
          <view v-if="memos.length === 0" class="empty-state">
            <text>暂无备忘录，请添加新的备忘录</text>
          </view>
          <view v-for="(memo, index) in memos" :key="index" class="memo-item" @tap="viewMemoDetail(memo, index)">
            <view class="memo-content">
              <text class="memo-title">{{ memo.title }}</text>
              <text class="memo-text">{{ memo.content }}</text>
              <text class="memo-date">{{ memo.date }}</text>
            </view>
            <view class="memo-actions">
              <text class="delete-btn" @tap.stop="deleteMemo(index)">🗑️</text>
            </view>
          </view>
        </scroll-view>
        <view class="memo-form">
          <input v-model="newMemo.title" placeholder="标题" class="memo-input" />
          <textarea v-model="newMemo.content" placeholder="内容" class="memo-textarea"></textarea>
          <button @tap="addMemo" class="add-btn">添加备忘录</button>
        </view>
        
      </view>

      <!-- 计时器模块 -->
      <view class="module timer-module">
        <view class="module-header">
          <text class="module-icon">⏰</text>
          <text class="module-title">计时器</text>
        </view>
        
        <!-- 场景选择 -->
        <view class="timer-scenes">
          <view v-for="scene in timerScenes" :key="scene.key" 
                :class="['scene-btn', { active: currentScene === scene.key }]"
                @tap="selectScene(scene.key)">
            <text>{{ scene.icon }} {{ scene.name }}</text>
          </view>
        </view>
        
        <!-- 主题选择 -->
        <view class="theme-selector">
          <text class="theme-label">主题：</text>
          <view v-for="theme in themes" :key="theme.key"
                :class="['theme-btn', { active: currentTheme === theme.key }]"
                :style="{ backgroundColor: theme.color }"
                @tap="selectTheme(theme.key)">
            <text :style="{ color: theme.textColor }">{{ theme.name }}</text>
          </view>
        </view>
        
        <!-- 计时器显示 -->
        <view class="timer-display" :style="{ backgroundColor: getCurrentTheme().color, color: getCurrentTheme().textColor }">
          <text class="timer-text">{{ formatTime(currentTime) }}</text>
        </view>
        
        <!-- 计时器控制 -->
        <view class="timer-controls">
          <view class="timer-mode">
            <view :class="['mode-btn', { active: timerMode === 'stopwatch' }]" @tap="setTimerMode('stopwatch')">
              <text>正计时</text>
            </view>
            <view :class="['mode-btn', { active: timerMode === 'countdown' }]" @tap="setTimerMode('countdown')">
              <text>倒计时</text>
            </view>
          </view>
          
          <view v-if="timerMode === 'countdown'" class="countdown-inputs">
            <input v-model.number="countdownHours" type="number" placeholder="时" class="time-input" />
            <input v-model.number="countdownMinutes" type="number" placeholder="分" class="time-input" />
            <input v-model.number="countdownSeconds" type="number" placeholder="秒" class="time-input" />
          </view>
          
          <view class="control-buttons">
            <button @tap="startTimer" :disabled="isRunning" class="control-btn start-btn">开始</button>
            <button @tap="pauseTimer" :disabled="!isRunning" class="control-btn pause-btn">暂停</button>
            <button @tap="resetTimer" class="control-btn reset-btn">重置</button>
          </view>
        </view>
        
        <!-- 背景音乐控制 -->
        <view class="music-controls">
          <text class="music-label">背景音乐：</text>
          <view class="music-options">
            <view v-for="music in musicOptions" :key="music.key"
                  :class="['music-btn', { active: currentMusic === music.key }]"
                  @tap="selectMusic(music.key)">
              <text>{{ music.name }}</text>
            </view>
          </view>
          <view class="music-controls-btn">
            <button @tap="toggleMusic" class="music-toggle-btn">{{ isMusicPlaying ? '暂停音乐' : '播放音乐' }}</button>
          </view>
           <!-- 音量调节横条 -->
          <view class="volume-control">
            <text class="volume-label">音量：</text>
            <slider 
              :value="audioVolume" 
              min="0" 
              max="1" 
              step="0.1" 
              show-value 
              @change="adjustVolume"
              class="volume-slider"
            />
          </view>
        </view>
        
        <!-- 提醒设置 -->
        <view class="reminder-settings">
          <text class="reminder-label">提醒设置：</text>
          <view class="reminder-options">
            <checkbox :checked="reminderEnabled" @tap="toggleReminder" />
            <text>启用提醒</text>
          </view>
          <input v-if="reminderEnabled" v-model="reminderMessage" placeholder="提醒消息" class="reminder-input" />
        </view>
      </view>

      <!-- 每日运势模块 -->
      <view class="module fortune-module">
        <view class="module-header">
          <text class="module-icon">🌟</text>
          <text class="module-title">每日运势</text>
        </view>
        <view class="fortune-content">
          <view class="fortune-card" :style="{ background: fortuneGradient }">
            <text class="fortune-date">{{ todayDate }}</text>
            <text class="fortune-icon">{{ fortuneIcon }}</text>
            <text class="fortune-text">{{ fortuneText }}</text>
            <text class="fortune-lucky">幸运数字: {{ luckyNumber }}</text>
          </view>
          <button @tap="drawFortune" :disabled="hasDrawnToday" class="fortune-btn">
            {{ hasDrawnToday ? '明天再来' : '抽取运势' }}
          </button>
        </view>
      </view>
    </view>
    
    <!-- 备忘录详情弹窗 -->
    <view v-if="showMemoDetail" class="modal-overlay" @tap="closeMemoDetail">
      <view class="memo-detail-modal" @tap.stop>
        <view class="modal-header">
          <text class="modal-title">备忘录详情</text>
          <text class="close-btn" @tap="closeMemoDetail">✕</text>
        </view>
        <view class="modal-content">
          <text class="detail-title">{{ selectedMemo.title }}</text>
          <text class="detail-content">{{ selectedMemo.content }}</text>
          <text class="detail-date">{{ selectedMemo.date }}</text>
        </view>
        <view class="modal-actions">
          <button @tap="deleteMemoFromDetail" class="delete-detail-btn">删除</button>
          <button @tap="closeMemoDetail" class="close-detail-btn">关闭</button>
        </view>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      username: '用户',
      // 备忘录相关
      memos: [],
      newMemo: {
        title: '',
        content: ''
      },
      showMemoDetail: false,
      selectedMemo: {},
      selectedMemoIndex: -1, 
      
      // 计时器相关
      currentTime: 0,
      isRunning: false,
      timerMode: 'stopwatch', // 'stopwatch' 或 'countdown'
      countdownHours: 0,
      countdownMinutes: 0,
      countdownSeconds: 0,
      timerInterval: null,
      currentScene: 'study',
      currentTheme: 'default',
      currentMusic: 'none',
      isMusicPlaying: false,
      reminderEnabled: false,
      reminderMessage: '时间到了！',
      audioVolume: 0.5,

      // 计时器场景
      timerScenes: [
        { key: 'study', name: '学习', icon: '📚' },
        { key: 'exercise', name: '健身', icon: '💪' },
        { key: 'exam', name: '考试', icon: '📝' },
        { key: 'pomodoro', name: '番茄钟', icon: '🍅' },
        { key: 'festival', name: '节日倒计时', icon: '🎉' }
      ],
      
      // 主题
      themes: [
        { key: 'default', name: '默认', color: '#6366f1', textColor: '#ffffff' },
        { key: 'green', name: '清新', color: '#10b981', textColor: '#ffffff' },
        { key: 'orange', name: '活力', color: '#f59e0b', textColor: '#ffffff' },
        { key: 'purple', name: '优雅', color: '#8b5cf6', textColor: '#ffffff' },
        { key: 'red', name: '热情', color: '#ef4444', textColor: '#ffffff' }
      ],
      
      // 音乐选项
      musicOptions: [
        { key: 'none', name: '无音乐' },
        { key: 'music1', name: '背景音乐1', src: '/static/music/background1.mp3' },
        { key: 'music2', name: '背景音乐2', src: '/static/music/background2.mp3' },
        { key: 'music3', name: '背景音乐3', src: '/static/music/background3.mp3' },
        { key: 'music4', name: '背景音乐4', src: '/static/music/background4.mp3' }
      ],
      bgAudio: null,
      
      // 运势相关
      fortuneText: '点击下方按钮抽取今日运势',
      fortuneIcon: '🔮',
      luckyNumber: 0,
      hasDrawnToday: false,
      fortuneGradient: 'linear-gradient(135deg, #6366f1, #8b5cf6)',
      
      // 运势数据
      fortunes: [
        { text: '今天的你充满活力，将遇到意想不到的好事。', icon: '☀️', gradient: 'linear-gradient(135deg, #f59e0b, #f97316)' },
        { text: '工作中可能遇到一些挑战，但坚持下去，终将成功。', icon: '⛰️', gradient: 'linear-gradient(135deg, #10b981, #059669)' },
        { text: '今天适合与朋友相聚，分享快乐时光。', icon: '👥', gradient: 'linear-gradient(135deg, #3b82f6, #2563eb)' },
        { text: '财运不错，可能有意外收获。', icon: '💰', gradient: 'linear-gradient(135deg, #f59e0b, #d97706)' },
        { text: '今天的你思维敏捷，是解决问题的好时机。', icon: '💡', gradient: 'linear-gradient(135deg, #6366f1, #4f46e5)' },
        { text: '健康状况良好，但也要注意休息。', icon: '❤️', gradient: 'linear-gradient(135deg, #ef4444, #dc2626)' },
        { text: '可能会收到一个好消息，让你心情愉悦。', icon: '📧', gradient: 'linear-gradient(135deg, #8b5cf6, #7c3aed)' },
        { text: '今天适合学习新知识，拓展视野。', icon: '📖', gradient: 'linear-gradient(135deg, #1d4ed8, #1e40af)' },
        { text: '人际关系和谐，沟通顺畅。', icon: '💬', gradient: 'linear-gradient(135deg, #10b981, #047857)' },
        { text: '创意灵感涌现，是开展创造性工作的好时机。', icon: '🎨', gradient: 'linear-gradient(135deg, #8b5cf6, #6d28d9)' }
      ]
    }
  },
  
  computed: {
    todayDate() {
      const now = new Date()
      return `${now.getFullYear()}年${now.getMonth() + 1}月${now.getDate()}日`
    }
  },
  
  onLoad() {
    this.loadMemos()
    this.checkFortuneStatus()
  },
  
  onUnload() {
    if (this.timerInterval) {
      clearInterval(this.timerInterval)
    }
  },
  
  methods: {
    // 备忘录相关方法
    loadMemos() {
      const currentUser = uni.getStorageSync('currentUser');
      if (currentUser) {
        const userMemos = uni.getStorageSync(`memos_${currentUser}`);
        this.memos = userMemos || [];
      }
    },
    
   adjustVolume(e) {
      this.audioVolume = e.detail.value;
      // 如果背景音乐正在播放，立即应用新的音量设置
      if (this.bgAudio && this.isMusicPlaying) {
        this.bgAudio.volume = this.audioVolume;
      }
    },

    saveMemos() {
      const currentUser = uni.getStorageSync('currentUser');
      if (currentUser) {
        uni.setStorageSync(`memos_${currentUser}`, this.memos);
      }
    },
    
    addMemo() {
      if (!this.newMemo.title.trim() || !this.newMemo.content.trim()) {
        uni.showToast({
          title: '请填写标题和内容',
          icon: 'none'
        })
        return
      }
      
      const now = new Date()
      const memo = {
        id: Date.now(),
        title: this.newMemo.title,
        content: this.newMemo.content,
        date: `${now.getFullYear()}-${(now.getMonth() + 1).toString().padStart(2, '0')}-${now.getDate().toString().padStart(2, '0')} ${now.getHours().toString().padStart(2, '0')}:${now.getMinutes().toString().padStart(2, '0')}`
      }
      
      this.memos.unshift(memo)
      this.saveMemos()
      
      this.newMemo.title = ''
      this.newMemo.content = ''
      
      uni.showToast({
        title: '添加成功',
        icon: 'success'
      })
    },
    
    deleteMemo(id) {
      // 阻止事件冒泡, Vue的@tap.stop会处理
      
      uni.showModal({
        title: '确认删除',
        content: '确定要删除这条备忘录吗？',
        success: (res) => {
          if (res.confirm) {
            this.memos = this.memos.filter(memo => memo.id !== id);
            this.saveMemos()
            uni.showToast({
              title: '删除成功',
              icon: 'success'
            })
          }
        }
      })
    },
    
    viewMemoDetail(memo, index) {
      this.selectedMemo = memo
      this.selectedMemoIndex = index
      this.showMemoDetail = true
    },
    
    closeMemoDetail() {
      this.showMemoDetail = false
      this.selectedMemo = {}
      this.selectedMemoIndex = -1
    },
    
    deleteMemoFromDetail() {
      this.deleteMemo(this.selectedMemo.id)
      this.closeMemoDetail()
    },
    
    // 计时器相关方法
    selectScene(sceneKey) {
      this.currentScene = sceneKey
      // 根据场景设置默认时间
      if (sceneKey === 'pomodoro') {
        this.countdownMinutes = 25
        this.timerMode = 'countdown'
      } else if (sceneKey === 'exercise') {
        this.countdownMinutes = 30
        this.timerMode = 'countdown'
      }
    },
    
    selectTheme(themeKey) {
      this.currentTheme = themeKey
    },
    
    getCurrentTheme() {
      return this.themes.find(theme => theme.key === this.currentTheme) || this.themes[0]
    },
    
        selectMusic(musicKey) {
      this.currentMusic = musicKey
      if (musicKey === 'none') {
        this.isMusicPlaying = false
        if (this.bgAudio) {
          this.bgAudio.stop()
        }
      } else if (this.isMusicPlaying && this.bgAudio) {
        // 如果正在播放音乐并切换了音乐选项，则应用新的音乐和音量设置
        const selectedMusic = this.musicOptions.find(music => music.key === musicKey)
        if (selectedMusic && selectedMusic.src) {
          this.bgAudio.src = selectedMusic.src
          this.bgAudio.volume = this.audioVolume; // 应用音量设置
          this.bgAudio.play()
        }
      }
    },
    
    toggleMusic() {
      if (this.currentMusic === 'none') {
        uni.showToast({
          title: '请先选择音乐',
          icon: 'none'
        })
        return
      }
      
      this.isMusicPlaying = !this.isMusicPlaying
      
      // 实际播放音乐
      const selectedMusic = this.musicOptions.find(music => music.key === this.currentMusic)
      
      if (this.isMusicPlaying) {
        if (!this.bgAudio) {
          this.bgAudio = uni.createInnerAudioContext()
          this.bgAudio.loop = true
        }
        
        if (selectedMusic && selectedMusic.src) {
          this.bgAudio.src = selectedMusic.src
          this.bgAudio.volume = this.audioVolume; // 应用音量设置
          this.bgAudio.play()
          
          uni.showToast({
            title: '音乐已播放',
            icon: 'none'
          })
        }
      } else if (this.bgAudio) {
        this.bgAudio.pause()
        
        uni.showToast({
          title: '音乐已暂停',
          icon: 'none'
        })
      }
    },
    
    toggleReminder() {
      this.reminderEnabled = !this.reminderEnabled
    },
    
    setTimerMode(mode) {
      this.timerMode = mode
      this.resetTimer()
    },
    
    startTimer() {
      if (this.timerMode === 'countdown') {
        const totalSeconds = this.countdownHours * 3600 + this.countdownMinutes * 60 + this.countdownSeconds
        if (totalSeconds <= 0) {
          uni.showToast({
            title: '请设置倒计时时间',
            icon: 'none'
          })
          return
        }
        this.currentTime = totalSeconds
      }
      
      this.isRunning = true
      this.timerInterval = setInterval(() => {
        if (this.timerMode === 'stopwatch') {
          this.currentTime++
        } else {
          this.currentTime--
          if (this.currentTime <= 0) {
            this.timerFinished()
          }
        }
      }, 1000)
    },
    
    pauseTimer() {
      this.isRunning = false
      if (this.timerInterval) {
        clearInterval(this.timerInterval)
        this.timerInterval = null
      }
    },
    
    resetTimer() {
      this.pauseTimer()
      this.currentTime = 0
    },
    
    timerFinished() {
      this.pauseTimer()
      
      if (this.reminderEnabled) {
        uni.showModal({
          title: '时间到了！',
          content: this.reminderMessage,
          showCancel: false
        })
      } else {
        uni.showToast({
          title: '时间到了！',
          icon: 'success'
        })
      }
      
      // 播放提醒音效
      uni.vibrateShort()
      
      // 播放提示音
      const alertAudio = uni.createInnerAudioContext()
      alertAudio.src = '/static/music/alert.mp3'
      alertAudio.play()
    },
    
    formatTime(seconds) {
      const hours = Math.floor(seconds / 3600)
      const minutes = Math.floor((seconds % 3600) / 60)
      const secs = seconds % 60
      return `${hours.toString().padStart(2, '0')}:${minutes.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`
    },
    
    // 运势相关方法
    checkFortuneStatus() {
      const today = new Date().toDateString()
      const lastDrawDate = uni.getStorageSync('fortuneDate')
      const lastFortune = uni.getStorageSync('lastFortune')
      
      if (lastDrawDate === today && lastFortune) {
        this.hasDrawnToday = true
        const fortune = JSON.parse(lastFortune)
        this.fortuneText = fortune.text
        this.fortuneIcon = fortune.icon
        this.luckyNumber = fortune.luckyNumber
        this.fortuneGradient = fortune.gradient
      }
    },
    
    drawFortune() {
      if (this.hasDrawnToday) {
        uni.showToast({
          title: '今天已经抽取过运势了',
          icon: 'none'
        })
        return
      }
      
      const randomIndex = Math.floor(Math.random() * this.fortunes.length)
      const fortune = this.fortunes[randomIndex]
      const luckyNumber = Math.floor(Math.random() * 100)
      
      this.fortuneText = fortune.text
      this.fortuneIcon = fortune.icon
      this.luckyNumber = luckyNumber
      this.fortuneGradient = fortune.gradient
      this.hasDrawnToday = true
      
      // 保存今日运势
      const today = new Date().toDateString()
      const fortuneData = {
        text: fortune.text,
        icon: fortune.icon,
        luckyNumber: luckyNumber,
        gradient: fortune.gradient
      }
      
      uni.setStorageSync('fortuneDate', today)
      uni.setStorageSync('lastFortune', JSON.stringify(fortuneData))
      
      uni.showToast({
        title: '运势抽取成功',
        icon: 'success'
      })
    }
  }
}
</script>

<style>
.dashboard {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 20rpx;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 30rpx;
  background: rgba(255, 255, 255, 0.9);
  border-radius: 20rpx;
  margin-bottom: 30rpx;
  box-shadow: 0 10rpx 30rpx rgba(0, 0, 0, 0.1);
}

.logo {
  display: flex;
  align-items: center;
}

.logo-img {
  width: 60rpx;
  height: 60rpx;
  margin-right: 20rpx;
}

.title {
  font-size: 36rpx;
  font-weight: bold;
  color: #333;
}

.user-info {
  display: flex;
  align-items: center;
}

.avatar {
  width: 60rpx;
  height: 60rpx;
  border-radius: 50%;
  background: linear-gradient(135deg, #6366f1, #8b5cf6);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-weight: bold;
  margin-right: 20rpx;
}

.username {
  font-size: 28rpx;
  color: #333;
}

.modules {
  display: flex;
  flex-direction: column;
  gap: 30rpx;
}

.module {
  background: rgba(255, 255, 255, 0.95);
  border-radius: 20rpx;
  padding: 30rpx;
  box-shadow: 0 10rpx 30rpx rgba(0, 0, 0, 0.1);
}

.module-header {
  display: flex;
  align-items: center;
  margin-bottom: 30rpx;
  padding-bottom: 20rpx;
  border-bottom: 2rpx solid #f0f0f0;
}

.module-icon {
  font-size: 40rpx;
  margin-right: 20rpx;
}

.module-title {
  font-size: 32rpx;
  font-weight: bold;
  color: #333;
}

/* 备忘录样式 */
.memo-list {
  max-height: 400rpx;
  margin-bottom: 30rpx;
}

.empty-state {
  text-align: center;
  color: #999;
  padding: 40rpx 0;
}

.memo-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20rpx;
  background: #f9f9f9;
  border-radius: 15rpx;
  margin-bottom: 20rpx;
  border-left: 6rpx solid #6366f1;
}

.memo-content {
  flex: 1;
}

.memo-title {
  font-size: 28rpx;
  font-weight: bold;
  color: #333;
  margin-bottom: 10rpx;
  display: block;
}

.memo-text {
  font-size: 24rpx;
  color: #666;
  margin-bottom: 10rpx;
  display: block;
}

.memo-date {
  font-size: 20rpx;
  color: #999;
  display: block;
}

.memo-actions {
  display: flex;
  align-items: center;
}

.delete-btn {
  font-size: 32rpx;
  color: #ef4444;
  padding: 10rpx;
}

.memo-form {
  display: flex;
  flex-direction: column;
  gap: 20rpx;
}

.memo-input, .memo-textarea {
  padding: 20rpx;
  border: 2rpx solid #e0e0e0;
  border-radius: 10rpx;
  font-size: 28rpx;
}

.memo-textarea {
  height: 120rpx;
}

.add-btn {
  padding: 20rpx;
  background: #6366f1;
  color: white;
  border: none;
  border-radius: 10rpx;
  font-size: 28rpx;
}

/* 计时器样式 */
.timer-scenes {
  display: flex;
  flex-wrap: wrap;
  gap: 15rpx;
  margin-bottom: 30rpx;
}

.scene-btn {
  padding: 15rpx 25rpx;
  background: #f0f0f0;
  border-radius: 25rpx;
  font-size: 24rpx;
  color: #666;
}

.scene-btn.active {
  background: #6366f1;
  color: white;
}

.theme-selector {
  display: flex;
  align-items: center;
  margin-bottom: 30rpx;
  flex-wrap: wrap;
  gap: 15rpx;
}

.theme-label {
  font-size: 28rpx;
  color: #333;
  margin-right: 20rpx;
}

.theme-btn {
  padding: 15rpx 25rpx;
  border-radius: 25rpx;
  font-size: 24rpx;
  border: 2rpx solid transparent;
}

.theme-btn.active {
  border-color: #333;
}

.timer-display {
  text-align: center;
  padding: 40rpx;
  border-radius: 20rpx;
  margin-bottom: 30rpx;
}

.timer-text {
  font-size: 60rpx;
  font-weight: bold;
  font-family: 'Courier New', monospace;
}

.timer-controls {
  display: flex;
  flex-direction: column;
  gap: 20rpx;
}

.timer-mode {
  display: flex;
  gap: 20rpx;
}

.mode-btn {
  flex: 1;
  padding: 20rpx;
  background: #f0f0f0;
  border-radius: 10rpx;
  text-align: center;
  font-size: 28rpx;
  color: #666;
}

.mode-btn.active {
  background: #6366f1;
  color: white;
}

.countdown-inputs {
  display: flex;
  gap: 20rpx;
}

.time-input {
  flex: 1;
  padding: 20rpx;
  border: 2rpx solid #e0e0e0;
  border-radius: 10rpx;
  text-align: center;
  font-size: 28rpx;
}

.control-buttons {
  display: flex;
  gap: 20rpx;
}

.control-btn {
  flex: 1;
  padding: 20rpx;
  border: none;
  border-radius: 10rpx;
  font-size: 28rpx;
  color: white;
}

.start-btn {
  background: #10b981;
}

.pause-btn {
  background: #f59e0b;
}

.reset-btn {
  background: #ef4444;
}

.control-btn:disabled {
  opacity: 0.5;
}

.music-controls, .reminder-settings {
  margin-top: 30rpx;
  padding-top: 30rpx;
  border-top: 2rpx solid #f0f0f0;
}

.music-label, .reminder-label {
  font-size: 28rpx;
  color: #333;
  margin-bottom: 20rpx;
  display: block;
}

.music-options {
  display: flex;
  flex-wrap: wrap;
  gap: 15rpx;
  margin-bottom: 20rpx;
}

.music-btn {
  padding: 15rpx 25rpx;
  background: #f0f0f0;
  border-radius: 25rpx;
  font-size: 24rpx;
  color: #666;
}

.music-btn.active {
  background: #8b5cf6;
  color: white;
}

.music-toggle-btn {
  padding: 20rpx 40rpx;
  background: #8b5cf6;
  color: white;
  border: none;
  border-radius: 10rpx;
  font-size: 28rpx;
}

.reminder-options {
  display: flex;
  align-items: center;
  gap: 15rpx;
  margin-bottom: 20rpx;
}

.reminder-input {
  padding: 20rpx;
  border: 2rpx solid #e0e0e0;
  border-radius: 10rpx;
  font-size: 28rpx;
  width: 100%;
}

/* 运势样式 */
.fortune-content {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.fortune-card {
  width: 100%;
  padding: 40rpx;
  border-radius: 20rpx;
  text-align: center;
  color: white;
  margin-bottom: 30rpx;
  position: relative;
  overflow: hidden;
}

.fortune-date {
  font-size: 24rpx;
  opacity: 0.8;
  margin-bottom: 20rpx;
  display: block;
}

.fortune-icon {
  font-size: 60rpx;
  margin-bottom: 30rpx;
  display: block;
}

.fortune-text {
  font-size: 32rpx;
  line-height: 1.6;
  margin-bottom: 30rpx;
  display: block;
}

.fortune-lucky {
  font-size: 24rpx;
  opacity: 0.8;
  display: block;
}

.fortune-btn {
  padding: 20rpx 40rpx;
  background: #6366f1;
  color: white;
  border: none;
  border-radius: 25rpx;
  font-size: 28rpx;
}

.fortune-btn:disabled {
  background: #ccc;
}

/* 弹窗样式 */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.memo-detail-modal {
  background: white;
  border-radius: 20rpx;
  padding: 40rpx;
  margin: 40rpx;
  max-width: 600rpx;
  width: 90%;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 30rpx;
  padding-bottom: 20rpx;
  border-bottom: 2rpx solid #f0f0f0;
}

.modal-title {
  font-size: 32rpx;
  font-weight: bold;
  color: #333;
}

.close-btn {
  font-size: 40rpx;
  color: #999;
  padding: 10rpx;
}

.modal-content {
  margin-bottom: 40rpx;
}

.detail-title {
  font-size: 32rpx;
  font-weight: bold;
  color: #333;
  margin-bottom: 20rpx;
  display: block;
}

.detail-content {
  font-size: 28rpx;
  color: #666;
  line-height: 1.6;
  margin-bottom: 20rpx;
  display: block;
}

.detail-date {
  font-size: 24rpx;
  color: #999;
  display: block;
}

.modal-actions {
  display: flex;
  gap: 20rpx;
}

.delete-detail-btn {
  flex: 1;
  padding: 20rpx;
  background: #ef4444;
  color: white;
  border: none;
  border-radius: 10rpx;
  font-size: 28rpx;
}

.close-detail-btn {
  flex: 1;
  padding: 20rpx;
  background: #6366f1;
  color: white;
  border: none;
  border-radius: 10rpx;
  font-size: 28rpx;
}

.volume-control {
  margin-top: 10px;
  padding: 8px;
  background-color: #f0f0f0;
  border-radius: 8px;
  display: flex;
  align-items: center;
}

.volume-label {
  font-size: 14px;
  margin-right: 10px;
  color: #333;
  min-width: 50px;
}

.volume-slider {
  flex: 1;
}

</style>