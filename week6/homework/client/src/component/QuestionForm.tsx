import React, { useEffect } from 'react';
import {
  Form,
  Input,
  Button,
  Space,
  Select,
  message,
  Checkbox,
  Radio,
} from 'antd';
import { saveQuestions } from '../api/questionApi';

const { Option } = Select;

const QuestionForm: React.FC<{
  initialData?: any;
  onCancel: () => void;
  onSuccess: () => void;
}> = ({ initialData, onCancel, onSuccess }) => {
  const [form] = Form.useForm();
  const type = Form.useWatch('type', form); //  监听题型变化
  const prevType = React.useRef<number | null>(null);
  const isInitialMount = React.useRef(true);

  useEffect(() => {
    if (initialData) {
      const {
        question,
        type,
        language,
        options = [],
        answer = [],
        ...rest
      } = initialData;
  
      form.setFieldsValue({
        ...rest,
        question,
        type,
        language,
        optionA: options[0],
        optionB: options[1],
        optionC: options[2],
        optionD: options[3],
        answer: answer,
      });
    }
  
    isInitialMount.current = false;
    prevType.current = initialData?.type ?? null;
  }, [initialData]);
  
  // 当题型变化时，清空与新类型不兼容的字段
  useEffect(() => {
    if (isInitialMount.current) return;
  
    if (type !== prevType.current) {
      if (type === 3) {
        // 切换到编程题，清空选项和答案
        form.setFieldsValue({
          optionA: undefined,
          optionB: undefined,
          optionC: undefined,
          optionD: undefined,
          answer: [],
        });
      } else if (type === 1 || type === 2) {
        // 切换为选择题，清空内容
        form.setFieldsValue({
          content: '',
        });
      }
    }
  
    prevType.current = type;
  }, [type]);

  const onFinish = async (values: any) => {
    const payload = {
      question: values.question,
      type: values.type,
      language: values.language,
      options:
        values.type === 3
          ? [] //  编程题不含选项
          : [values.optionA, values.optionB, values.optionC, values.optionD],
      answer: values.type === 3 ? [] : values.answer,
      explanation: values.explanation || '',
      content: values.content || '',
      aiStartTime: new Date().toISOString(),
      aiEndTime: new Date().toISOString(),
      aiCostTime: 0,
      aiReq: {
        model: values.model || '',
        language: values.language[0] || '',
        type: values.type,
        keyword: '',
        difficulty: values.difficulty || '',
        count: 1,
      },
      aiRes: {
        question: values.question,
        options:
          values.type === 3
            ? []
            : [values.optionA, values.optionB, values.optionC, values.optionD],
        answer: values.answer || [],
        explanation: values.explanation || '',
      },
    };

    try {
      await saveQuestions([payload]);
      message.success('保存成功');
      onSuccess();
    } catch (err) {
      message.error('保存失败');
    }
  };

  return (
    <Form
      form={form}
      layout="vertical"
      onFinish={onFinish}
      style={{ height: 700, overflowY: 'auto', padding: 0 }}
    >
      <Form.Item name="type" label="题型" rules={[{ required: true }]}>
        <Select placeholder="请选择题型">
          <Option value={1}>单选题</Option>
          <Option value={2}>多选题</Option>
          <Option value={3}>编程题</Option>
        </Select>
      </Form.Item>

      <Form.Item name="language" label="语言" rules={[{ required: true }]}>
        <Select mode="multiple" placeholder="请选择语言">
          <Option value="go语言">Go语言</Option>
          <Option value="java">Java</Option>
          <Option value="python">Python</Option>
        </Select>
      </Form.Item>

      <Form.Item name="question" label="标题" rules={[{ required: true }]}>
        <Input />
      </Form.Item>

      <Form.Item name="content" label="内容">
        <Input.TextArea rows={2} showCount maxLength={500} />
      </Form.Item>

      {/* 选项和答案仅在非编程题时显示 */}
      {type !== 3 && (
        <>
          <Form.Item label="选项 A" name="optionA" rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <Form.Item label="选项 B" name="optionB" rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <Form.Item label="选项 C" name="optionC" rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <Form.Item label="选项 D" name="optionD" rules={[{ required: true }]}>
            <Input />
          </Form.Item>

          {type === 1 && (
            <Form.Item
              name="answer"
              label="答案"
              rules={[{ required: true, message: '请选择一个答案' }]}
            >
              <Radio.Group>
                <Space direction="horizontal">
                  <Radio value="A">A</Radio>
                  <Radio value="B">B</Radio>
                  <Radio value="C">C</Radio>
                  <Radio value="D">D</Radio>
                </Space>
              </Radio.Group>
            </Form.Item>
          )}

          {type === 2 && (
            <Form.Item
              name="answer"
              label="答案"
              rules={[{ required: true, message: '请选择至少一个答案' }]}
            >
              <Checkbox.Group>
                <Space direction="horizontal">
                  <Checkbox value="A">A</Checkbox>
                  <Checkbox value="B">B</Checkbox>
                  <Checkbox value="C">C</Checkbox>
                  <Checkbox value="D">D</Checkbox>
                </Space>
              </Checkbox.Group>
            </Form.Item>
          )}

          <Form.Item name="explanation" label="解析">
            <Input.TextArea rows={2} showCount maxLength={500} />
          </Form.Item>
        </>
      )}


      <Form.Item>
        <Space>
          <Button onClick={onCancel}>取消</Button>
          <Button type="primary" htmlType="submit">
            保存
          </Button>
        </Space>
      </Form.Item>
    </Form>
  );
};

export default QuestionForm;
