// 性能数据

export function logPerformanceMetrics() {
  // 白屏时间
  const timing = performance.timing
  const whiteScreenTime = timing.responseStart - timing.navigationStart

  // 首屏时间
  const paintEntries = performance.getEntriesByType('paint')
  const firstContentfulPaint = paintEntries.find(entry => entry.name === 'first-contentful-paint')

  // 可交互时间
  const domReadyTime = timing.domContentLoadedEventEnd - timing.navigationStart

  // 完全加载时间
  const loadTime = timing.loadEventEnd - timing.navigationStart

  // 打印结果
  console.log('White Screen Time:', whiteScreenTime)
  if (firstContentfulPaint) {
    console.log('First Contentful Paint Time:', firstContentfulPaint.startTime)
  } else {
    console.log('First Contentful Paint data not available.')
  }
  console.log('Time to Interactive:', domReadyTime)
  console.log('Load Time:', loadTime)
}

