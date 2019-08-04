desc 'Format source codes'
task :fmt do
  sh 'find `pwd` -type d -name lib -prune -o -type f -name "*.go" | grep -v vendor | xargs -t -n 1 goimports -w '
end

desc 'Lint'
task :lint do
  `glide novendor -x`.split.each do |target|
    sh "golint -set_exit_status #{target} || exit $?"
  end
end

desc 'Build binary'
task :build do
  sh 'git rev-parse --is-inside-work-tree' do |ok, status|
    if ok
      version = `git describe --tags --abbrev=0`.chomp
      revision = `git rev-parse --short HEAD`.chomp
    else
      version = '0.0'
      revision = 'xxxxxxxx'
    end

    ldflags = "-X main.version=#{version} -X main.revision=#{revision}"

    sh "go build -ldflags \"#{ldflags}\" main.go"
  end
end
