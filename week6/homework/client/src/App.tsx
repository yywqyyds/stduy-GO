import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate, useLocation, useNavigate } from 'react-router-dom';
import QuestionManager from './pages/QuestionManager';
import { StudyPage } from './pages/StudyPage';

import { Layout, Menu } from 'antd';

const { Header, Content, Sider } = Layout;

// 提取 SiderMenu 组件，负责处理跳转和高亮
const SiderMenu: React.FC = () => {
  const navigate = useNavigate();
  const location = useLocation();

  const items = [
    { key: '/study', label: '学习页面' },
    { key: '/questions', label: '题库管理' }
  ];

  return (
    <Menu
      mode="inline"
      selectedKeys={[location.pathname]} // 高亮当前菜单
      style={{ height: '100%' }}
      items={items}
      onClick={(e) => navigate(e.key)} // 点击跳转
    />
  );
};

const App: React.FC = () => {
  return (
    <Router>
      <Layout style={{ minHeight: '100vh' }}>
      <Header style={{ display: 'flex', alignItems: 'center', background: '#001529' }}>
        <img src="public/logo.png" alt="logo" style={{ height: 32, marginRight: 12 }} />
        <span style={{ color: '#fff', fontSize: 18 }}>武汉科技大学 宋浩沅大作业</span>
      </Header>
        <Layout>
          <Sider width={200}>
            <SiderMenu />
          </Sider>
          <Layout>
            <Content style={{ padding: '24px' }}>
              <Routes>
                <Route path="/" element={<Navigate to="/questions" />} />
                <Route path="/study" element={<StudyPage />} />
                <Route path="/questions" element={<QuestionManager />} />
              </Routes>
            </Content>
          </Layout>
        </Layout>
      </Layout>
    </Router>
  );
};

export default App;
