import React, { useState, useRef } from 'react';
import { Form, Input, Select, Button, message, Row, Col, Card, Spin } from 'antd';
import { createQuestions, saveQuestions } from '../api/questionApi';
import { useNavigate } from 'react-router-dom';


const { Option } = Select;

const AiQuestionGenerator: React.FC = () => {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState(false);
  const [generatedQuestions, setGeneratedQuestions] = useState<any[]>([]);

  const navigate = useNavigate();

  const generatedQuestionsRef = useRef<any[]>([]);
  
  const handleGenerate = async () => {
    try {
      const values = await form.validateFields();
      setLoading(true);
  
      const res = await createQuestions(values);
  
      if (res.data.code === 0) {
        generatedQuestionsRef.current = res.data.data;
        setGeneratedQuestions(res.data.data);
        console.log('保存用数据:', generatedQuestionsRef.current);
      }
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  };
  

  return (
    <Row gutter={24}>
      {/* 左侧出题表单 */}
      <Col span={10}>
        <Form form={form} layout="vertical" initialValues={{ type: 1, count: 1, language: 'go语言', difficulty: '简单' }}>
          <Form.Item label="题目类型" name="type" rules={[{ required: true }]}>
            <Select>
              <Option value={1}>单选题</Option>
              <Option value={2}>多选题</Option>
              <Option value={3}>编程题</Option>
            </Select>
          </Form.Item>

          <Form.Item label="生成数量" name="count" rules={[{ required: true }]}>
            <Select>
                {Array.from({ length: 10 }, (_, i) => (
                  <Option key={i + 1} value={i + 1}>
                    {i + 1}
                  </Option>
                ))}
              </Select>
          </Form.Item>

          <Form.Item label="语言" name="language" rules={[{ required: true }]}>
            <Select>
              <Option value="go">go语言</Option>
              <Option value="javascript">js语言</Option>
            </Select>
          </Form.Item>

          <Form.Item label="关键词" name="keyword" >
                  <Input placeholder='请输入关键词，如gin,路由等' />
          </Form.Item>

          <Form.Item label="难度" name="difficulty" rules={[{ required: true }]}>
            <Select>
              <Option value="简单">简单</Option>
              <Option value="中等">中等</Option>
              <Option value="困难">困难</Option>
            </Select>
          </Form.Item>

          <Form.Item>
            <Button type="primary" onClick={handleGenerate} loading={loading}>
              生成
            </Button>
          </Form.Item>
        </Form>
      </Col>

      {/* 右侧预览展示 */}
      <Col span={14}>
        <div style={{ maxHeight: '400px', overflowY: 'auto', paddingRight: 8 }}>
          <Spin spinning={loading}>
            {generatedQuestions.length === 0 ? (
              <p style={{ marginTop: 20, color: '#999' }}>暂无生成结果</p>
            ) : (
              <>
                {generatedQuestions.map((item, index) => (
                <Card key={index} title={`题目 ${index + 1}`} style={{ marginBottom: 16 }}>
                  <p><strong>题目：</strong>{item.aiRes.question}</p>

                  {/* 如果是选择题，显示选项 */}
                  {item.aiRes.options?.length > 0 && (
                    <>
                      <p><strong>选项：</strong></p>
                      <ul>
                        {item.aiRes.options.map((opt: string, idx: number) => (
                          <li key={idx}>{opt}</li>
                        ))}
                      </ul>
                    </>
                  )}

                  {/* 仅当 answer 存在并非空数组时显示 */}
                  {Array.isArray(item.aiRes.answer) && item.aiRes.answer.length > 0 && (
                    <p><strong>答案：</strong>{item.aiRes.answer.join(', ')}</p>
                  )}

                  {/* 仅当 explanation 存在且非空字符串时显示 */}
                  {item.aiRes.explanation && item.aiRes.explanation.trim() !== '' && (
                    <p><strong>解析：</strong>{item.aiRes.explanation}</p>
                  )}
                </Card>
              ))}


                {/* 保存按钮放到底部 */}
                <Button
                  type="primary"
                  style={{ marginTop: 16 }}
                  onClick={async () => {
                    try {
                      const res = await saveQuestions(generatedQuestionsRef.current);
                      console.log('保存响应结果:', res);
                      
                      const status = res.status
                      if (status === 200) {
                        message.success('保存成功');
                        setTimeout(() => {
                          navigate('/'); // 你可以替换为主界面的具体路径
                        }, 1000);
                      } else {
                        message.error('保存失败');
                      }
                    } catch (err) {
                      console.error('保存请求异常:', err);
                      message.error('保存失败');
                    }
                  }}
                >
                  保存题目
                </Button>

              </>
            )}
          </Spin>
        </div>
      </Col>
    </Row>
  );
};

export default AiQuestionGenerator;
