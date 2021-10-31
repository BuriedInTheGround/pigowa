{
  description = "Go interface for using p5.js in the browser.";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      with nixpkgs.legacyPackages.${system}; rec {
        packages.pigowa-mouse-events = stdenv.mkDerivation rec {
          pname = "pigowa-mouse-events";
          version = "0.0.0";

          name = "${pname}-${version}";

          src = fetchFromGitHub {
            owner = "BuriedInTheGround";
            repo = "pigowa";
            rev = "c8b9ad8fe22fd40757f8f69b683732c0c94549fc";
            sha256 = "sha256-UxKJdy4r4zv1MubR/2c7/Gw9wsahiykgIJVDGyBnZTw=";
          };

          meta = with lib; {
            description = "Go interface for using p5.js in the browser.";
            homepage = "https://github.com/BuriedInTheGround/pigowa";
            license = licenses.bsd3;
          };

          buildInputs = [ go_1_16 ];

          exampleToBuild = "mouse-events";

          configurePhase = ''
            runHook preConfigure

            export GOCACHE=$TMPDIR/go-cache
            export GOPATH="$TMPDIR/go"
            export GOSUMDB=off

            runHook postConfigure
          '';

          buildPhase = ''
            runHook preBuild

            cd ./examples/${exampleToBuild}
            GOOS=js GOARCH=wasm go build -o main.wasm main.go
            cd ../../

            go install ./server

            runHook postBuild
          '';

          installPhase = ''
            runHook preInstall

            mkdir -p $out/${exampleToBuild}

            cd ./examples/${exampleToBuild}
            cp index.html $out/${exampleToBuild}
            cp wasm_exec.js $out/${exampleToBuild}
            cp main.wasm $out/${exampleToBuild}

            dir="$GOPATH/bin"
            [ -e "$dir" ] && cp -r $dir $out

            runHook postInstall
          '';
        };

        devShell = mkShell {
          buildInputs = [ go_1_16 ];
        };
      }
    );
}
