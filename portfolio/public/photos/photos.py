import os
import json

def list_photos(directory):
    photos = []
    for filename in os.listdir(directory):
        if filename.lower().endswith(('.png', '.jpg', '.jpeg', '.gif', '.bmp')):
            photos.append(filename)
    return photos

if __name__ == "__main__":
    directory = os.path.dirname(os.path.abspath(__file__))
    photos = list_photos(directory)
    output_file = os.path.join(directory, 'photos.json')
    with open(output_file, 'w') as f:
        json.dump(photos, f, indent=4)