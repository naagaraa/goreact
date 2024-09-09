import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { ConfigService } from '@nestjs/config';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  const configService = app.get(ConfigService);
  const port = configService.get<number>('NEST_PORT') || 3000; // Default to port 3000 if not specified
  await app.listen(port);
  console.log(`Application is running on: http://localhost:${port}`);
}
bootstrap();
