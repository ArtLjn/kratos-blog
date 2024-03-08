export const utilButton = () => {
    const startPosition = window.pageYOffset; // 获取当前滚动位置
    const targetPosition = 850; // 目标滚动位置

    const distance = targetPosition - startPosition; // 计算滚动距离
    const duration = 1000; // 滚动持续时间（以毫秒为单位）
    const startTime = performance.now(); // 获取动画开始时间

    const scrollAnimation = (currentTime) => {
      const elapsedTime = currentTime - startTime; // 计算已过去的时间
      const progress = Math.min(elapsedTime / duration, 1); // 计算滚动进度（0到1之间）

      const easeProgress = easeOutQuad(progress); // 使用缓动函数调整滚动进度（可根据需要更改缓动函数）

      const newPosition = startPosition + distance * easeProgress; // 计算新的滚动位置
      window.scrollTo(0, newPosition); // 滚动到新位置

      if (elapsedTime < duration) {
        // 如果动画未完成，继续执行动画
        window.requestAnimationFrame(scrollAnimation);
      }
    };

    // 缓动函数 - 使用了二次缓动曲线
    const easeOutQuad = (t) => {
      return t * (2 - t);
    };

    // 开始滚动动画
    window.requestAnimationFrame(scrollAnimation);
  };