import os
import requests
import json
from minio import Minio
headers= {
    "Cookie": "ory_kratos_session=MTYxOTQzNDE4OXxEdi1CQkFFQ180SUFBUkFCRUFBQVJfLUNBQUVHYzNSeWFXNW5EQThBRFhObGMzTnBiMjVmZEc5clpXNEdjM1J5YVc1bkRDSUFJR0YyZUVGeGRuQjFlVmhVYkdFelZqQmxjRUpsVTNOYVIyeEJWVXhNVGtaSnzHsZiXwB4kQ3HDYGCSPYARQsPVQU6SITrU86ARwi8egg==",
    "X-Space": "1"
}
path = "./templates/"
client = Minio(
    endpoint="127.0.0.1:9000",
    access_key="miniokey",
    secret_key="miniosecret",
    secure=False,
)
def create_medium(path, chart, filename):
    found = client.bucket_exists("dega")
    if found:
        print("Bucket 'dega' already exists")
    res = client.fput_object(
        "dega", "/bindu/"+ chart, os.path.abspath(path+filename), "image/png"
    )
    print("Successfully uploaded")

    body = {
        "name": chart,
        "url": {
            "raw" : "http://localhost:9000/dega/bindu/" + chart
        }
    }
    res = requests.post(
        url= "http://127.0.0.1:4455/.factly/bindu/server/media",
        json=body,
        headers=headers
    )
    print(res.text)
    return res

categories = os.listdir(path)
categories.append(categories.pop(categories.index('Others')))

while len(categories) > 0:
    category = categories.pop(0)
    category_path = path + category + "/"
    dirs = os.listdir(category_path)
    print(f'category: {category}')
    body = {
        "name": category,
    }
    res = requests.post(
        url= "http://127.0.0.1:4455/.factly/bindu/server/categories",
        json=body,
        headers=headers
    )
    print (res.json())
    data = res.json()

    category_id = data['id']
    while len(dirs) > 0:
        dir = dirs.pop(0)
        if os.path.isdir(category_path + dir):
            dir_path = category_path + dir + "/"
            files = os.listdir(dir_path)
            for file in files:
                file_path = dir_path + file
                if file == "spec.json":
                    f = open(file_path, "r")
                    spec = json.loads(f.read())
                    f.close()
                elif file == "properties.json":
                    f = open(file_path, "r")
                    properties = json.loads(f.read())    
                    f.close()
            res = create_medium(dir_path,dir + ".png" , "thumbnail.png")
            print(res.json())
            data = res.json()
            print(f'chart: {dir}')
            body = {
                "category_id": category_id,
                "medium_id": data['id'],
                "properties": properties,
                "slug": dir.replace(" ", "-"),
                "spec": spec,
                "title": dir,
            }
            requests.post(
                url= "http://127.0.0.1:4455/.factly/bindu/server/templates",
                json=body,
                headers=headers
            )
    print(f'\n')