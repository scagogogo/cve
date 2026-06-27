import { ConfigProvider, theme } from 'antd'
import zhCN from 'antd/locale/zh_CN'
import HomePage from './pages/HomePage'

function App() {
  return (
    <ConfigProvider
      locale={zhCN}
      theme={{
        algorithm: theme.defaultAlgorithm,
        token: {
          colorPrimary: '#e94560',
          borderRadius: 8,
          fontFamily: '-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif',
        },
      }}
    >
      <HomePage />
    </ConfigProvider>
  )
}

export default App
