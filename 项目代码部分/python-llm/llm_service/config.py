"""
LLM服务配置
"""
import os
from typing import Optional
from pydantic_settings import BaseSettings


class LLMSettings(BaseSettings):
    """大模型配置"""
    # 模型提供商: openai, anthropic, qianfan
    provider: str = "openai"

    # 模型名称
    model: str = "gpt-4"

    # API密钥
    api_key: Optional[str] = None

    # API基础URL
    base_url: Optional[str] = None

    # 温度参数
    temperature: float = 0.7

    # 最大Token数
    max_tokens: int = 2000

    # 请求超时（秒）
    timeout: int = 60


class RAGSettings(BaseSettings):
    """RAG配置"""
    # 向量数据库类型: chroma, milvus
    vector_db: str = "chroma"

    # 向量数据库路径
    persist_directory: str = "./data/vector_db"

    # 嵌入模型
    embedding_model: str = "sentence-transformers/all-MiniLM-L6-v2"

    # 检索Top-K
    top_k: int = 3

    # 相似度阈值
    similarity_threshold: float = 0.7


class ServerSettings(BaseSettings):
    """服务配置"""
    # 主机地址
    host: str = "0.0.0.0"

    # 端口
    port: int = 8000

    # 日志级别
    log_level: str = "info"


class Settings:
    """全局配置"""
    llm: LLMSettings = LLMSettings()
    rag: RAGSettings = RAGSettings()
    server: ServerSettings = ServerSettings()

    class Config:
        env_file = ".env"
        env_file_encoding = "utf-8"


# 全局配置实例
settings = Settings()
