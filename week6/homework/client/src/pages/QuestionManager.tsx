import React, { useEffect, useState } from 'react';
import {
  Table,
  Button,
  Input,
  Modal,
  message,
  Dropdown,
  Menu,
  Tag,
} from 'antd';
import { DownOutlined } from '@ant-design/icons';
import type { ColumnsType } from 'antd/es/table';
import { getQuestionList, deleteQuestions, queryQuestionById } from '../api/questionApi';
import QuestionForm from '../component/QuestionForm';
import AiQuestionGenerator from '../component/AiQuestionGenerator';

const { Search } = Input;

interface Question {
  id: number;
  question: string;
  type: number;
  difficulty?: string;
  creator?: string;
}

const TYPE_MAP: Record<string, number | undefined> = {
  '全部': undefined,
  '单选': 1,
  '多选': 2,
  '编程': 3,
};

const TYPE_LABEL: Record<number, string> = {
  1: '单选',
  2: '多选',
  3: '编程',
};

const QuestionManager: React.FC = () => {
  const [questions, setQuestions] = useState<Question[]>([]);
  const [loading, setLoading] = useState(false);
  const [typeFilter, setTypeFilter] = useState<string>('全部');
  const [searchText, setSearchText] = useState('');
  const [selectedRowKeys, setSelectedRowKeys] = useState<React.Key[]>([]);
  const [aiModalVisible, setAiModalVisible] = useState(false);


  const fetchData = async (keywordParam?: string) => {
    const mappedType = TYPE_MAP[typeFilter];
    const keyword = keywordParam !== undefined ? keywordParam : searchText;

    const params: { type?: number; keyword?: string } = {};
    if (mappedType !== undefined) params.type = mappedType;
    if (keyword.trim()) params.keyword = keyword.trim();

    setLoading(true);
    try {
      const res = await getQuestionList(params);
      const { code, data } = res.data;
      if (code === 0 && Array.isArray(data)) {
        setQuestions(data);
      } else {
        message.warning('题目列表返回异常');
      }
    } catch (err) {
      message.error('加载失败');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, [typeFilter]);

  const handleDelete = async (id: number) => {
    Modal.confirm({
      title: '确认删除该题目吗？',
      onOk: async () => {
        try {
          await deleteQuestions([id]);
          message.success('删除成功');
          fetchData();
        } catch (err) {
          message.error('删除失败');
        }
      },
    });
  };

  const handleBatchDelete = async () => {
    Modal.confirm({
      title: `确认删除选中的 ${selectedRowKeys.length} 个题目吗？`,
      onOk: async () => {
        await deleteQuestions(selectedRowKeys as number[]);
        message.success('批量删除成功');
        setSelectedRowKeys([]);
        fetchData();
      },
    });
  };

  const [editingQuestion, setEditingQuestion] = useState<Question | null>(null);
  const [editModalVisible, setEditModalVisible] = useState(false);

  const handleEdit = async (id: number) => {
    try {
      const res = await queryQuestionById(id);
      const { data } = res.data;
      const processed = {
        id,
        question: data.aiRes.question,
        type: data.aiReq.type,
        language: [data.aiReq.language],
        options: data.aiRes.options,
        answer: data.aiRes.answer,
        explanation: data.aiRes.explanation,
        content: '',
      };
      setEditingQuestion(processed);
      setEditModalVisible(true);
    } catch (err) {
      message.error('获取题目失败');
      console.error(err);
    }
  };

  const dropdownMenu = (
    <Menu
      onClick={({ key }) => {
        if (key === 'ai') {
          setAiModalVisible(true); // 显示 AI 出题弹窗
        } else {
          setEditingQuestion(null);
          setEditModalVisible(true); // 人工出题弹窗
        }
      }}
    >
      <Menu.Item key="ai">AI 出题</Menu.Item>
      <Menu.Item key="manual">人工出题</Menu.Item>
    </Menu>
  );

  const columns: ColumnsType<Question> = [
    {
      title: '题目',
      dataIndex: 'question',
    },
    {
      title: '题型',
      dataIndex: 'type',
      render: (type) => <Tag color={type === 1 ? 'blue' : type === 2 ? 'green' : 'orange'}>{TYPE_LABEL[type] || type}</Tag>,
    },
    {
      title: '操作',
      width: 160,
      render: (_, record) => (
        <>
          <Button type="link" onClick={() => handleEdit(record.id)}>
            编辑
          </Button>
          <Button type="link" danger onClick={() => handleDelete(record.id)}>
            删除
          </Button>
        </>
      ),
    },
  ];

  return (
    <div>
      <div style={{ marginBottom: 16, display: 'flex', alignItems: 'center', gap: 16 }}>
        <span style={{ fontWeight: 500 }}>题型：</span>
        {Object.keys(TYPE_MAP).map((key) => (
          <Button
            key={key}
            type={typeFilter === key ? 'primary' : 'default'}
            shape="round"
            onClick={() => setTypeFilter(key)}
          >
            {key}
          </Button>
        ))}

        <Search
          placeholder="请输入试题名称"
          allowClear
          style={{ width: 200 }}
          onSearch={(value) => {
            setSearchText(value);
            fetchData(value);
          }}
        />

        <Button danger disabled={selectedRowKeys.length === 0} onClick={handleBatchDelete}>
          批量删除
        </Button>

        <Dropdown overlay={dropdownMenu}>
          <Button type="primary">
            + 出题 <DownOutlined />
          </Button>
        </Dropdown>
      </div>

      <Table
        rowKey="id"
        loading={loading}
        columns={columns}
        dataSource={questions}
        rowSelection={{
          selectedRowKeys,
          onChange: (keys) => setSelectedRowKeys(keys),
        }}
        pagination={{ pageSize: 10 }}
      />

      <Modal
        title="AI 出题"
        open={aiModalVisible}
        onCancel={() => setAiModalVisible(false)}
        footer={null}
        width={800}
      >
        <AiQuestionGenerator />
      </Modal>

      <Modal
        title={editingQuestion ? '编辑题目' : '新增题目'}
        open={editModalVisible}
        onCancel={() => setEditModalVisible(false)}
        footer={null}
      >
        <QuestionForm
          initialData={editingQuestion}
          onCancel={() => setEditModalVisible(false)}
          onSuccess={() => {
            setEditModalVisible(false);
            fetchData();
          }}
        />
      </Modal>
    </div>
  );
};

export default QuestionManager;
