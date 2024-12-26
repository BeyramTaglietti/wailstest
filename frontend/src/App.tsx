import { useEffect, useState } from "react";
import { ReadFile, ToBlackAndWhite } from "../wailsjs/go/main/App";
import { OnFileDrop, OnFileDropOff } from "../wailsjs/runtime";
import { Button } from "./components/ui/button";

function App() {
  const [base64Image, setBase64Image] = useState<string | null>(null);
  const [originalImagePath, setOriginalImagePath] = useState<string | null>(
    null
  );

  const [transformedImage, setTransformedImage] = useState<string | null>(null);

  useEffect(() => {
    OnFileDrop((_, __, paths) => {
      const path = paths[0];
      setOriginalImagePath(path);

      ReadFile(path).then((data) => {
        setBase64Image(data);
      });
    }, true);

    return () => {
      OnFileDropOff();
    };
  }, []);

  const handleTransformImage = () => {
    if (!originalImagePath) {
      return;
    }

    ToBlackAndWhite(originalImagePath).then((data) => {
      ReadFile(data).then((data) => {
        setTransformedImage(data);
      });
    });
  };

  return (
    <>
      <div className="flex flex-col h-dvh w-dvw p-4 gap-4">
        <div className="flex flex-row h-1/2 w-full gap-4">
          <div
            className="border rounded overflow-hidden w-1/2"
            style={{ "--wails-drop-target": "drop" } as React.CSSProperties}
          >
            {base64Image && (
              <img
                src={`data:image/png;base64,${base64Image}`}
                className="w-full h-full"
                alt="image"
              />
            )}
          </div>
          <div className="border rounded overflow-hidden w-1/2">
            {transformedImage && (
              <img
                src={`data:image/png;base64,${transformedImage}`}
                className="w-full h-full"
                alt="transformed image"
              />
            )}
          </div>
        </div>

        <Button onClick={handleTransformImage}>Transform image</Button>
      </div>
    </>
  );
}

export default App;
