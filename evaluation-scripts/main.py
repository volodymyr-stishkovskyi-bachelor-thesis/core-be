from pinecone import Pinecone
import openai
import csv
import numpy as np

openai.api_key = ""

# 2) Pinecone client (unchanged)
pc = Pinecone(
    api_key="",
    environment="gcp-starter",
)
ix = pc.Index("")


def embed(text: str) -> list[float]:
    resp = openai.embeddings.create(model="text-embedding-ada-002", input=text)
    return resp.data[0].embedding


def precision_and_mrr(csv_path: str) -> tuple[float, float]:
    p_at_3 = 0
    reciprocal_ranks = 0
    n = 0

    with open(csv_path, newline="") as f:
        reader = csv.DictReader(f)
        for row in reader:
            # 4) Use keyword args in Pinecone query
            vec = embed(row["query"])
            res = ix.query(vector=vec, top_k=10, include_metadata=False)
            ids = [match["id"] for match in res["matches"]]

            # Precision@3
            relevant = set(v for v in row.values() if v)
            hits3 = any(doc_id in relevant for doc_id in ids[:3])

            # Reciprocal Rank
            rr = 0.0
            for rank, doc_id in enumerate(ids, start=1):
                if doc_id in relevant:
                    rr = 1.0 / rank
                    break

            p_at_3 += hits3
            reciprocal_ranks += rr
            n += 1

    return p_at_3 / n, reciprocal_ranks / n


if __name__ == "__main__":
    p, mrr = precision_and_mrr("gold.csv")
    print(f"P@3={p:.2f}, MRR={mrr:.2f}")


# P@3=0.47, MRR=0.45
