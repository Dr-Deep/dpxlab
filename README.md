# README



# dpxlab.de nginx
* Poudriere build status
* pkg repos (Cloudflare buckets) 
* radicle http mirror
* Impressum (about me)



# BDE: Build Decision Engine
* relevanter commit (branch art?, letzder build? wegen cooldown)

# Build Engine
* commit reference nehmen
* GCP Build VM starten
* artefakte auf Cloudflare R2 bewegen



Build Engine Controller







GO + GCP API


# dpxlab Build Engine Config

# dpxlab Build Engine (lib, wie gcp vm und so)
* checkout src / ports
* make / poudriere bulk
* poudriere repo
=> Cloudflare R2

# dpxlab Build Engine Controller
* was hat sich geändert / relevant? / 
* branch?
* letzder build aufgrund cooldown / grade ein build schon aktiv?
* commits auf pkg makefile versionierung fokussieren?
* build-dep: src -> port/port-update ; sodass nie inkonsistent ist

# dpxlab Build Artefact Storage
* Cloudflare R2 via S3 API 
* via http/html/css schmücken vllcht
* url: "https://pkg.dpxlab.de/${ABI}/latest" -> MANIFEST


# dpxlab Build States (SQL Tables)
* builds: status, version, commit/id, start/stop time, logs

# ToDo
* nginx

# ToBe
* commit hook (src,ports)
* 


Nginx
Storage buckets für pkg repo
The heart
Radicle git server mit github mirror / github bot


## src hook
* radicle mirror hook
* compilen für ein image
* cloudflare buckets



## ports hook


# make.conf / src.conf / poudriere.conf
* CCACHE
* -O3
* ThinLTO
* LLVM-BOLT
* Vector Flags: -fopt-info-vec -falign-functions=32
* -ffunction-sections -fdata-sections
* LDFLAGS+= -Wl,--gc-sections -Wl,--icf=all
* -mllvm -inline-threshold=350 -mllvm -unroll-threshold=150
* -mllvm -polly -mllvm -polly-vectorizer=stripmine
* -Wl,--lto-O3 ?


Port und src building infra für custom cflags


Terraform emacs language server
Terraform cli
Github repo + deploy
Cloudflare buckets für pkg repo storage unter dpxlab.de

