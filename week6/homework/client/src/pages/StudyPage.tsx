import React, { useEffect, useState } from 'react';
import ReactMarkdown from 'react-markdown';
import { Spin, Alert } from 'antd';

export const StudyPage: React.FC = () => {
  const [markdown, setMarkdown] = useState<string>('');
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetch('../public/docs/readme.md')
      .then((res) => {
        if (!res.ok) {
          throw new Error(`HTTP error ${res.status}`);
        }
        return res.text();
      })
      .then((text) => {
        setMarkdown(text);
        setLoading(false);
      })
      .catch((err) => {
        setError(err.message);
        setLoading(false);
      });
  }, []);

  if (loading) return <Spin tip="加载中..." />;
  if (error) return <Alert type="error" message="加载失败" description={error} />;

  return (
    <div style={{ background: '#fff', padding: '24px', borderRadius: '6px' }}>
      <ReactMarkdown>{markdown}</ReactMarkdown>
    </div>
  );
};
