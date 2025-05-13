import axios from 'axios';

const baseURL = '/api/questions';

// 调用大模型出题（不保存，仅预览）
export const createQuestions = async (data: {
  model?: string;
  language?: string;
  type: number;
  keyword: string;
  difficulty: string;
  count: number;
}) => {
  return axios.post(`${baseURL}/create`, data);
};

// 获取题库列表，支持按类型和关键词搜索
export const getQuestionList = async (params?: {
  type?: number;
  keyword?: string;
}) => {
  const query: any = {};
  if (params?.type) {
    query.type = params.type;
  }
  if (params?.keyword) {
    query.keyword = params.keyword;
  }
  return axios.get('/api/questions/list', { params: query });
};

// 查询题目详情（按 ID 查询）
export const queryQuestionById = async (id: number) => {
  return axios.get(`${baseURL}/query/${id}`);
};

// 删除题目（支持批量删除）
export const deleteQuestions = async (ids: number[]) => {
  return axios.post(`${baseURL}/delete`, { ids });
};

// 保存题目（支持单题或多题保存）
export const saveQuestions = async (questions: {
  aiStartTime: string;
  aiEndTime: string;
  aiCostTime: number;
  aiReq: {
    model: string;
    language: string;
    type: number;
    keyword: string;
    count: number;
    difficulty: string;
  };
  aiRes: {
    question: string;
    options: string[];
    answer: string[];
    explanation: string;
  };
}[]) => {
  return axios.post('/api/questions/save', questions);
};
